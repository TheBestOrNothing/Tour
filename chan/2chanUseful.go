package main

import (
	"fmt"
)

func main() {
	done := make(chan bool)
	go func() {
		fmt.Println("go routinge message")
		//... do some thing ...
		//... do some thing ...
		//... do some thing ...
		done <- true
	}()

	<-done
	fmt.Println("Main function message")
}
