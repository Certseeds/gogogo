package mapreduce

import (
	"errors"
	"fmt"
	"log"
	"math"
	"net"
	"net/http"
	"sync"
	"time"
)
import "os"
import "net/rpc"

type Coordinator struct {
	// Your definitions here.
}

const (
	kBegin      = 1
	kMapping    = 2
	kMapDone    = 3
	kReducing   = 4
	kReduceDone = 5
	kNotReach   = 0xFFFFFFFF
)

// workerMap type should be <string,int>
var (
	workerMap         = make(map[int]workerStatus)
	mapOrderNumber    = make(map[string]int) // readonly except MakeCoordinator
	mapTaskMap        = make(map[string]int) // 0 for not run, 1 for running, 2 for done
	reduceTaskMap     = make(map[int]int)    // 0 for not run, 1 for running, 2 for done
	canReturn         = false
	workerSignChannel = make(chan SignRequest)
	workerSignDone    = make(chan bool)
	workerReadAll     = make(chan struct{})
	workerReadAllDone = make(chan map[int]workerStatus)
	workerUpdateAll   = make(chan map[int]workerStatus)
	workerUpdateDone  = make(chan struct{})
	canReturnInput    = make(chan bool)
	canReturnOutput   = make(chan bool)
)

func init() {
	go func() {
		for {
			select {
			case request := <-workerSignChannel:
				{
					workerAddress := request.SockNum
					if _, exist := workerMap[workerAddress]; exist {
						workerSignDone <- false
						continue
					}
					workerMap[workerAddress] = workerStatus{
						enable:     true,
						mapTask:    "",
						reduceTask: 0,
					}
					workerSignDone <- true
				}
			case <-workerReadAll:
				{
					willReturn := make(map[int]workerStatus)
					for k, v := range workerMap {
						if v.enable {
							willReturn[k] = v
						}
					}
					workerReadAllDone <- willReturn
				}
			case update := <-workerUpdateAll:
				{
					for k, v := range update {
						workerMap[k] = v
					}
					workerUpdateDone <- struct{}{}
				}
			}
		}
	}()
	go func() {
		for {
			select {
			case input := <-canReturnInput:
				{
					canReturn = input
				}
			case canReturnOutput <- canReturn:
			}
		}
	}()
}

type workerStatus struct {
	enable     bool
	mapTask    string
	reduceTask int
}

// Your code here -- RPC handlers for the worker to call.

func (c *Coordinator) Sign(request SignRequest, reply *EmptyResponse) error {
	workerAddress := request.SockNum
	workerSignChannel <- request
	result := <-workerSignDone
	if !result {
		return errors.New("address exist")
	}
	log.Println(workerAddress)
	return nil
}

// Done main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
func (c *Coordinator) Done() bool {
	value := <-canReturnOutput
	log.Printf("isDone? %v", value)
	return value
}

type workerDone struct {
	num    int
	status int
}

func (c *Coordinator) cycle() {
	for {
		// key is rpc:port, value is status
		log.Println("before get from global")
		workerReadAll <- struct{}{}
		workers := <-workerReadAllDone
		log.Println("after get from global")
		func(workers_ *map[int]workerStatus) {
			disableChannel := make(chan int)
			doneChannel := make(chan workerDone)
			passChannel := make(chan struct{})
			wg := &sync.WaitGroup{}
			for port, _ := range *workers_ {
				wg.Add(1)
				go watchWorkerStatus(port, disableChannel, doneChannel, passChannel, wg)
			}
			updateWorkerStatus(disableChannel, doneChannel, passChannel, len(*workers_), workers_)
			wg.Wait()
		}(&workers)
		log.Println("done get from workers")
		// then, in all key-value in workerMap if value == 0, then can put tasks
		// if value > 0 then should ask for status
		status := judgeingStatus()
		log.Printf("%d", status)
		switch status {
		case kBegin:
			fallthrough
		case kMapping:
			mapSchedule(&workers)
		case kMapDone:
			fallthrough
		case kReducing:
			reduceSchedule(&workers)
		case kReduceDone:
			beforeExit(&workers)
			return
		case kNotReach:
			fallthrough
		default:
			log.Println("should not have this state of coordinator")
		}
		log.Println("before update to global")
		workerUpdateAll <- workers
		<-workerUpdateDone
		log.Println("after update to global")
		time.Sleep(time.Second)
	}
}
func mapSchedule(status *map[int]workerStatus) {
	todoFiles := make([]string, 0)
	for FileName, statu := range mapTaskMap {
		if statu == 0 {
			todoFiles = append(todoFiles, FileName)
		}
	}
	type pair struct {
		key   int
		value workerStatus
	}
	runnableWorker := make([]pair, 0)
	for port, statu := range *status {
		if len(statu.mapTask) == 0 && len(todoFiles) > 0 {
			runnableWorker = append(runnableWorker, pair{port, statu})
		}
	}
	minLength := int(math.Min(float64(len(todoFiles)), float64(len(runnableWorker))))
	statusChan := make(chan pair, minLength)
	for i := 0; i < minLength; i++ {
		go func(todoFile string, port int, pairChan chan pair) {
			callWorker(port, "WorkWork.Mapper", MapRequest{
				FileName:   todoFile,
				MapOrder:   mapOrderNumber[todoFile],
				ReduceNums: len(reduceTaskMap),
			}, nil)
			pairChan <- pair{
				key: port,
				value: workerStatus{
					enable:     true,
					mapTask:    todoFile,
					reduceTask: 0,
				},
			}
		}(todoFiles[i], runnableWorker[i].key, statusChan)
	}
	func(length int, pairChan <-chan pair, status_ *map[int]workerStatus) {
		for count := 0; count < length; {
			select {
			case p := <-pairChan:
				{
					mapTaskMap[p.value.mapTask] = 1
					(*status_)[p.key] = p.value
					count += 1
				}
			}
		}
	}(minLength, statusChan, status)
}
func reduceSchedule(status *map[int]workerStatus) {
	todoTasks := make([]int, 0)
	for taskOrder, statu := range reduceTaskMap {
		if statu == 0 {
			todoTasks = append(todoTasks, taskOrder+1)
		}
	}
	if len(todoTasks) > 0 {
		// TODO, use goroutine replace for cycle
		for port, statu := range *status {
			if statu.reduceTask == 0 && len(todoTasks) > 0 {
				todoTask := todoTasks[0]
				todoTasks = todoTasks[1:]
				reduceFileList := make([]string, len(mapTaskMap))
				fileNamer := func(mapTaskNum int) string {
					return fmt.Sprintf("./mr-%d-%d", mapTaskNum, todoTask)
				}
				for i := 0; i < len(mapTaskMap); i++ {
					reduceFileList[i] = fileNamer(i)
				}
				callWorker(port, "WorkWork.Reducer", ReduceRequest{
					ReduceFileList: reduceFileList,
					ReduceOrder:    todoTask,
				}, nil)
				reduceTaskMap[todoTask-1] = 1
				(*status)[port] = workerStatus{
					enable:     true,
					mapTask:    "",
					reduceTask: todoTask,
				}
			}
		}
	}
}

func beforeExit(status *map[int]workerStatus) {
	// send exit to each workers
	wg := &sync.WaitGroup{}
	for port_, _ := range *status {
		wg.Add(1)
		go func(port int, wg1 *sync.WaitGroup) {
			defer wg.Done()
			callWorker(port, "WorkWork.Exit", struct{}{}, nil)
		}(port_, wg)
	}
	wg.Wait()
	canReturnInput <- true
}
func judgeingStatus() int {
	allMapTaskUndo := func() bool {
		for _, state := range mapTaskMap {
			if state != 0 {
				return false
			}
		}
		return true
	}()
	allReduceTaskUndo := func() bool {
		for _, state := range reduceTaskMap {
			if state != 0 {
				return false
			}
		}
		return true
	}()
	if allMapTaskUndo && allReduceTaskUndo {
		// all zeros
		return kBegin
	}
	// then, one of them should have an unzero element
	allMapTaskDone := func() bool {
		for _, state := range mapTaskMap {
			if state != 2 {
				return false
			}
		}
		return true
	}()
	allReduceTaskDone := func() bool {
		for _, state := range reduceTaskMap {
			if state != 2 {
				return false
			}
		}
		return true
	}()
	if allMapTaskDone && allReduceTaskDone {
		// all 2
		return kReduceDone
	} else if allMapTaskDone && allReduceTaskUndo {
		// all 2 and all 0
		return kMapDone
	}
	for _, state := range reduceTaskMap {
		if state != 0 {
			return kReducing
		}
	}
	for _, state := range mapTaskMap {
		if state != 0 {
			return kMapping
		}
	}
	return kNotReach
}

func updateWorkerStatus(disableChannel chan int, doneChannel chan workerDone, passChannel chan struct{}, max int, localMap *map[int]workerStatus) {
	for count := 0; count < max; {
		select {
		case num := <-disableChannel:
			{
				workStatus := (*localMap)[num]
				workStatus.enable = false
				if len(workStatus.mapTask) != 0 {
					mapTaskMap[workStatus.mapTask] = 0
					workStatus.mapTask = ""
				} else if workStatus.reduceTask != 0 {
					reduceTaskMap[workStatus.reduceTask-1] = 0
					workStatus.reduceTask = 0
				}
				(*localMap)[num] = workStatus
				count += 1
				log.Printf("disable %d %d\n", num, count)
			}
		case done := <-doneChannel:
			{
				workStatus := (*localMap)[done.num]
				if done.status == kWorkerMapDone {
					mapTaskMap[workStatus.mapTask] = 2
					workStatus.mapTask = ""
				} else if done.status == kWorkerReduceDone {
					reduceTaskMap[workStatus.reduceTask-1] = 2
					workStatus.reduceTask = 0
				}
				(*localMap)[done.num] = workStatus
				count += 1
				log.Printf("disable %v %d\n", done, count)
			}
		case <-passChannel:
			{
				count += 1
				log.Printf("pass %d\n", count)
			}
		}
	}

}
func watchWorkerStatus(num int, disableChannel chan int, doneChannel chan workerDone, passChannel chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	reply := &AliveResponse{DoneOrDoing: 0, TaskStartTime: 0}
	ok := callWorker(num, "WorkWork.Alive", AliveRequest{Secret: AreYouOK}, reply)
	log.Println(num, ok, *reply)
	if !ok {
		disableChannel <- num
	} else if reply.DoneOrDoing == kWorkerMapDone || reply.DoneOrDoing == kWorkerReduceDone {
		doneChannel <- workerDone{num: num, status: reply.DoneOrDoing}
	} else if reply.DoneOrDoing == kWorkerMapDoing || reply.DoneOrDoing == kWorkerReduceDoing {
		timeDiff := time.Now().UnixNano() - reply.TaskStartTime
		if timeDiff >= 10*time.Second.Nanoseconds() {
			disableChannel <- num
		} else {
			passChannel <- struct{}{}
		}
	} else {
		passChannel <- struct{}{}
	}
	return
}

// MakeCoordinator create a Coordinator.
// main/mrcoordinator.go calls this function.
// nReduce is the number of reduce tasks to use.
func MakeCoordinator(files []string, nReduce int) *Coordinator {
	// ps: the most simple way is: each file have a map task
	c := Coordinator{}
	c.server(coordinatorSock(os.Getuid()))
	// Your code here.
	for order, file := range files {
		mapTaskMap[file] = 0
		mapOrderNumber[file] = order
	}
	for i := 0; i < nReduce; i++ {
		reduceTaskMap[i] = 0
	}
	go c.cycle()
	return &c
}

func callWorker(addr int, rpcname string, args interface{}, reply interface{}) bool {
	c, err := rpc.DialHTTP("unix", coordinatorSock(addr))
	if err != nil {
		log.Print("dialing:", err)
		return false
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	log.Println(err)
	return false
}

// start a thread that listens for RPCs from worker.go
func (c *Coordinator) server(sockname string) {
	rpc.Register(c)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}
