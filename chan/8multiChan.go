package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	TimeFrom = 1
	TimeTo   = 3
)

func timeOut() int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(TimeTo-TimeFrom) + TimeFrom
	return randNum
}
func getMsgChan() chan string {
	c := make(chan string)
	rand := timeOut()
	time.Sleep(time.Duration(rand) * time.Second)
	go func() {
		c <- fmt.Sprintf("hello: %d", rand)
	}()
	return c
}
func main() {
	c1 := getMsgChan()
	c2 := getMsgChan()
	c3 := getMsgChan()

	//ok1, ok2, ok3 := true, true, true
	for {
		select {
		case msg, ok := <-c1:
			fmt.Println(msg, ok)
			if ok {
				c1 = nil
			}
		case msg, ok := <-c2:
			fmt.Println(msg, ok)
			if ok {
				c2 = nil
			}
		case msg, ok := <-c3:
			fmt.Println(msg, ok)
			if ok {
				c3 = nil
			}
		}
		if c1 == nil && c2 == nil && c3 == nil {
			fmt.Println("all chan close and return main")
			return
		}
	}

	fmt.Println("Main function message")
}
