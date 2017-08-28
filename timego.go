package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	TimeFrom   = 1
	TimeTo     = 2
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

func resetTimer(t *time.Timer, d time.Duration) {
	stopTimer(t)
	t.Reset(d)
}

func main() {
	timeoutT := time.NewTimer(time.Duration(1) * time.Second)
	go func() {
		for {
			select {
			case <-timeoutT.C:
				fmt.Println("Timeout and stratup election")
			default:
			}
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			resetTimer(timeoutT, time.Duration(1)*time.Second)
		}
	}()

	go func() {
		time.Sleep(time.Second * 8)
		stopTimer(timeoutT)
	}()

	for {

		/*
			go func() {
				time.Sleep(time.Second * 10)
				stopTimer(timeoutT)
			}()
		*/
	}

}
