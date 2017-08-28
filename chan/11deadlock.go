package main

import "fmt"

func main() {
	//writing operations on a synchronous channel can only succeed
	//in case there is a receiver ready to get this data.
	//And we create the receiver only in the next line.
	c := make(chan int, 1)
	c <- 42    // write to a channel
	val := <-c // read from a channel
	fmt.Println(val)
}
