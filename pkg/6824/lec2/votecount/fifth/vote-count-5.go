package main

import (
	"gogogo/pkg/6824/lec2/votecount"
	"time"
)
import "math/rand"

func main() {
	rand.Seed(time.Now().UnixNano())

	count := 0
	finished := 0
	ch := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			ch <- votecount.RequestVote()
		}()
	}
	for count < 5 && finished < 10 {
		v := <-ch
		if v {
			count += 1
		}
		finished += 1
	}
	if count >= 5 {
		println("received 5+ votes!")
	} else {
		println("lost")
	}
}
