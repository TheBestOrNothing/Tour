package main

import (
	"fmt"
	"time"
)

func funcWithUnpredictiveExecutionTime() bool {
	time.Sleep(100 * time.Millisecond)
	return true
}
func measureTime(expectedMs float64) {
	done := make(chan bool)
	t1 := time.Now()

	go func() {
		done <- funcWithUnpredictiveExecutionTime()
		//close(done)
	}()

	for {
		fmt.Println("for for ....")
		select {
		case state := <-done:
			fmt.Printf("Done state:%v\n", state)
			//break
			goto END
		case <-time.After(time.Duration(expectedMs) * time.Millisecond):
			fmt.Println("Do another thing ....")
		}
	}
END:

	fmt.Println(time.Since(t1))
	return

}

func main() {
	//measureTime(1000)
	measureTime(10)
}
