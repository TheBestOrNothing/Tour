package main

import (
	"fmt"
	"time"
)

type Raft struct {
	reset chan bool
}

func (rf *Raft) NewTimer() {

	timeout := make(chan bool)
	go func() {
		time.Sleep(time.Duration(200) * time.Millisecond)
		timeout <- true
	}()

	for {
		fmt.Printf("NewTimer ..  \n")
		select {
		case <-rf.reset:
			go rf.NewTimer()
			return
		case <-timeout:
			fmt.Printf("f2c..  \n")
			//f2c(rf)
			//default:
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
