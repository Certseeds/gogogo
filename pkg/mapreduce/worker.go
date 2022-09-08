package mapreduce

import (
	"fmt"
)
import "hash/fnv"

// Map functions return a slice of KeyValue.
type KeyValue struct {
	Key   string
	Value string
}

// use ihash(key) % NReduce to choose the reduce
// task number for each KeyValue emitted by Map.
func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

// main/mrworker.go calls this function.
func Worker(mapf func(string, string) []KeyValue,
	reducef func(string, []string) string) {

	// Your worker implementation here.

	// uncomment to send the MainExample RPC to the coordinator.
	CallExample()

}

// example function to show how to make an RPC call to the coordinator.
//
// the RPC argument and reply types are defined in rpc.go.
func CallExample() {

	// declare an argument structure.
	args := ExampleArgs{}

	// fill in the argument(s).
	args.X = 99

	// declare a reply structure.
	reply := ExampleReply{}

	// send the RPC request, wait for the reply.
	// the "Coordinator.MainExample" tells the
	// receiving server that we'd like to call
	// the MainExample() method of struct Coordinator.
	ok := call("Coordinator.MainExample", args, &reply)

	if ok {
		// reply.Y should be 100.
		fmt.Printf("reply.Y %v \n", reply.Y)
	} else {
		fmt.Printf("call failed!\n")
	}
}
