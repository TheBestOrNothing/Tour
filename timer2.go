package main

import (
	"fmt"
	"time"
)

type Raft struct {
	reset chan bool
}

func (rf *Raft) NewTimer() {

	fmt.Printf("NewTimer ..  \n")
	timeout := make(chan bool, 1)

	go func() {
		rf.Sleep(time.Duration(200) * time.Millisecond)
		timeout <- true
	}()

	for {
		select {
		case <-rf.reset:
			go rf.NewTimer()
			return
		case <-timeout:
			fmt.Println("timeout")
			break
		}
	}
}

func (rf *Raft) Sleep(d time.Duration) {
	t0 := time.Now()
	for time.Since(t0) < d {
	}
}

func (rf *Raft) ResetTimer() {
	i := 0
	for {
		fmt.Printf("ReSetTimer   %d\n", i)
		rf.Sleep(time.Duration(50) * time.Millisecond)
		go func() { rf.reset <- true }()
		i++
	}
}

func main() {
	// ... do some things, maybe even return under some condition
	rf := &Raft{}
	rf.reset = make(chan bool)
	go rf.NewTimer()
	rf.ResetTimer()
	return
}
