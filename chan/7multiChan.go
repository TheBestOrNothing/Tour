package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	TimeFrom = 1
	TimeTo   = 5
)

func timeOut() int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(TimeTo-TimeFrom) + TimeFrom
	return randNum
}
func getMsgChan() <-chan string {
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

	for {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		case msg := <-c3:
			fmt.Println(msg)
		}
	}

	fmt.Println("Main function message")
}
