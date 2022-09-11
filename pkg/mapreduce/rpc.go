package mapreduce

//
// RPC definitions.
//
// remember to capitalize all names.
//

import (
	"fmt"
)

// example to show how to declare the arguments
// and reply for an RPC.
type SignRequest struct {
	SockNum int
}
type EmptyResponse struct{}

const AreYouOK = "Kept you waiting, huh?"

type AliveRequest struct {
	Secret string
}

const (
	kWorkerWaiting     = 1
	kWorkerMapDoing    = 2
	kWorkerMapDone     = 3
	kWorkerReduceDoing = 4
	kWorkerReduceDone  = 5
)

type AliveResponse struct {
	DoneOrDoing   int   // 1 for waiting, 2 for doing, 3 for Done
	TaskStartTime int64 // the unixstyle-timestamp for start the task
}

const BaseMagic = 100000007

type MapRequest struct {
	FileName   string
	MapOrder   int // when input, add BasicMagic , reciver then minus it
	ReduceNums int // when input, add BasicMagic , reciver then minus it
}
type ReduceRequest struct {
	ReduceFileList []string
	ReduceOrder    int
}

// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the coordinator.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func coordinatorSock(num int) string {
	s := fmt.Sprintf("/var/tmp/824-mr-%d", num)
	return s
}
