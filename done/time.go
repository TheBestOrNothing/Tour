package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	TimeFrom   = 1
	TimeTo     = 6
	HeartBeats = 6 * time.Millisecond
	Leader     = 0
	Candidate  = 1
	Follower   = 2
)

func timeOut() time.Duration {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(TimeTo-TimeFrom) + TimeFrom
	return time.Duration(randNum) * time.Second
}

func stopTimer(t *time.Timer) {
	if !t.Stop() {
		select {
		case <-t.C:
			fmt.Println("resetTimer: What happend??")
			fmt.Println("resetTimer: This means timeout happend but no electioin trigged!")
		default:
		}
	}
}

func reset(t *time.Timer) {
	go func() { resetTimer(t) }()
}

func main() {
	timeoutD := timeOut()
	timeoutT := time.NewTimer(timeoutD)
	for {
		time.Sleep(time.Second)
		select {
		case <-timeoutT.C:
			fmt.Println("Timeout and stratup election")
			resetTimer(timeoutT, timeOut())
		default:
			fmt.Println("No time out")
		}
	}
}
