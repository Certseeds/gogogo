package mapreduce

import (
	"bufio"
	"errors"
	"fmt"
	"hash/fnv"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"sort"
	"sync/atomic"
	"time"
)

type WorkWork struct {
}

var (
	doingStatus         atomic.Int32
	doingBeginTimestamp atomic.Int64
	portGlobal          int
)

func init() {
	doingStatus.Store(kWorkerWaiting)
}

// KeyValue Map functions return a slice of KeyValue.
type KeyValue struct {
	Key   string
	Value string
}
type KVList []KeyValue

func (s KVList) Len() int {
	return len(s)
}
func (s KVList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s KVList) Less(i, j int) bool {
	return s[i].Key < s[j].Key
}

var mapFunction func(string, string) []KeyValue
var reduceFunction func(string, []string) string

// use ihash(key) % NReduce to choose the reduce
// task number for each KeyValue emitted by Map.
func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

func (c *WorkWork) Mapper(request_ MapRequest, response *EmptyResponse) error {
	fileNamer := func(reduceTaskNum int) string {
		return fmt.Sprintf("./mr-%d-%d", request_.MapOrder, reduceTaskNum)
	}
	doingStatus.Store(kWorkerMapDoing)
	doingBeginTimestamp.Store(time.Now().UnixNano())
	go func(request MapRequest) {
		log.Println(portGlobal, request.FileName, request.MapOrder, request.ReduceNums)
		defer doingStatus.Store(kWorkerMapDone)
		fileList := make([][]KeyValue, request.ReduceNums)
		for i := 0; i < request.ReduceNums; i++ {
			fileList[i] = make([]KeyValue, 0)
		}
		file, _ := os.Open(request.FileName)
		content, _ := io.ReadAll(file)
		file.Close()
		output := mapFunction(request.FileName, string(content))
		for _, kv := range output {
			hash := ihash(kv.Key) % request.ReduceNums
			fileList[hash] = append(fileList[hash], kv)
		}
		for order, kvList := range fileList {
			fileName := fileNamer(order + 1)
			file2, _ := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
			//
			write := bufio.NewWriter(file2)
			for _, kv := range kvList {
				write.WriteString(fmt.Sprintf("%s %s\n", kv.Key, kv.Value))
			}
			write.Flush()
			file2.Close()
		}
	}(request_)
	return nil
}
func readFileByLine(fileName string) []string {
	file, _ := os.Open(fileName)
	defer file.Close()
	willReturn := make([]string, 0)
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		willReturn = append(willReturn, scanner.Text())
	}
	return willReturn
}
func (c *WorkWork) Reducer(request_ ReduceRequest, response *EmptyResponse) error {
	fileList := request_.ReduceFileList
	doingStatus.Store(kWorkerReduceDoing)
	go func(request ReduceRequest) {
		log.Println(portGlobal, request.ReduceOrder, request.ReduceFileList)
		defer doingStatus.Store(kWorkerReduceDone)
		allLines := make([]string, 0)
		for _, filePath := range fileList {
			lines := readFileByLine(filePath)
			allLines = append(allLines, lines...)
		}
		log.Println(portGlobal, request.ReduceOrder, "read Reduce Files ", len(allLines))
		allKVPair := make([]KeyValue, 0)
		for _, line := range allLines {
			var key string
			var value string
			fmt.Sscanf(line, "%v %v", &key, &value)
			allKVPair = append(allKVPair, KeyValue{Key: key, Value: value})
		}
		log.Println(portGlobal, request.ReduceOrder, "read kv pairs", len(allKVPair))
		sort.Sort(KVList(allKVPair))
		log.Println(portGlobal, request.ReduceOrder, "sort by keys", len(allKVPair))
		keyMapToValueList := make(map[string][]string, 0)
		for i, lastStr := 0, ""; i < len(allKVPair); i++ {
			if allKVPair[i].Key != lastStr {
				lastStr = allKVPair[i].Key
				keyMapToValueList[lastStr] = make([]string, 0)
			}
			keyMapToValueList[lastStr] = append(keyMapToValueList[lastStr], allKVPair[i].Value)
		}
		log.Println(portGlobal, request.ReduceOrder, "divide by keys", len(keyMapToValueList))
		reduceResultList := make([]string, 0)
		for key, values := range keyMapToValueList {
			output := reduceFunction(key, values)
			reduceResultList = append(reduceResultList, fmt.Sprintf("%v %v\n", key, output))
		}
		defer log.Println(portGlobal, request.ReduceOrder, "write out done", len(reduceResultList))
		file, _ := os.OpenFile(fmt.Sprintf("mr-out-%d", request_.ReduceOrder), os.O_CREATE|os.O_RDWR, 0644)
		defer file.Close()
		write := bufio.NewWriter(file)
		defer write.Flush()
		for _, line := range reduceResultList {
			write.WriteString(line)
		}
	}(request_)
	return nil
}
func (c *WorkWork) Exit(request_ struct{}, response *EmptyResponse) error {
	os.Exit(0)
	return nil
}

// Worker main/mrworker.go calls this function.
func Worker(mapf func(string, string) []KeyValue, reducef func(string, []string) string) {
	rand.Seed(time.Now().UnixNano())
	mapFunction = mapf
	reduceFunction = reducef
	// Your worker implementation here.

	// uncomment to send the MainExample RPC to the coordinator.
	for count := 0; count < 3; count += 1 {
		if Sign() {
			break
		}
	}
	select {}
}

// Sign example function to show how to make an RPC call to the coordinator.
//
// the RPC argument and reply types are defined in rpc.go.
func Sign() bool {
	addrPort := 1024 + rand.Intn(4096)
	ok := callCoor("Coordinator.Sign", SignRequest{SockNum: addrPort}, nil)
	if !ok {
		return false
	}
	addr := coordinatorSock(addrPort)
	portGlobal = addrPort
	ww := WorkWork{}
	ww.server(addr)
	return true
}

// start a thread that listens for RPCs from worker.go
func (c *WorkWork) server(sockname string) {
	rpc.Register(c)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go func() {
		defer os.Remove(sockname)
		http.Serve(l, nil)
	}()
}
func (c *WorkWork) Alive(request AliveRequest, reply *AliveResponse) error {
	if request.Secret != AreYouOK {
		return errors.New("snake? snake, snake!! ")
	}
	reply.DoneOrDoing = int(doingStatus.Load())
	if reply.DoneOrDoing == kWorkerMapDoing || reply.DoneOrDoing == kWorkerReduceDoing {
		reply.TaskStartTime = doingBeginTimestamp.Load()
	} else if reply.DoneOrDoing == kWorkerMapDone || reply.DoneOrDoing == kWorkerReduceDone {
		reply.TaskStartTime = time.Now().UnixNano()
	}
	return nil
}

func callCoor(rpcname string, args interface{}, reply interface{}) bool {
	// c, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	sockname := coordinatorSock(os.Getuid())
	c, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	log.Println(err)
	return false
}
