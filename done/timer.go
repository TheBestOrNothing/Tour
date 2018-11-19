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
	timeout := make(chan bool)
	defer func() {
		time.Sleep(time.Duration(201) * time.Millisecond)
		select {
		case <-timeout:
		default:
		}
		fmt.Printf("XXXTimer ..  \n")
	}()

	go func() {
		time.Sleep(time.Duration(200) * time.Millisecond)
		timeout <- true
	}()

	for {
		//DPrintf("ForTimer1 ..  %d(T:%2d)(S:%d)\n", rf.me, rf.currentTerm, rf.status)
		select {
		case <-rf.reset:
			go rf.NewTimer()
			return
		case <-timeout:
			fmt.Printf("f2c..  \n")
			break
		}
	}

}

func (rf *Raft) ResetTimer() {
	i := 0
	for {
		fmt.Printf("ReSetTimer   %d\n", i)
		time.Sleep(time.Duration(50) * time.Millisecond)
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
