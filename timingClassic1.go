package main

import (
	"fmt"
	"time"
)

func funcWithUnpredictiveExecutionTime() {
	time.Sleep(500 * time.Millisecond)
}
func measureTime(expectedMs float64) {
	done := make(chan struct{})
	t1 := time.Now()

	go func() {
		funcWithUnpredictiveExecutionTime()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Done ")
	case <-time.After(time.Duration(expectedMs) * time.Millisecond):
		fmt.Println("Do another thing")
	}

	fmt.Println(time.Since(t1))
	return

}

func main() {
	measureTime(1000)
}
