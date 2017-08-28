package main

import (
	"fmt"
)

func main() {
	msg := make(chan string)
	count := 3
	go func() {
		fmt.Println("go routinge message")
		for idx := 0; idx < count; idx++ {
			//... do some thing ...
			//... do some thing ...
			//... do some thing ...
			msg <- fmt.Sprintf("hello%d", idx)
		}
		close(msg)
	}()

	for m := range msg {
		fmt.Println(m)
	}
	fmt.Println(<-msg)
	fmt.Println(<-msg)
	fmt.Println("Main function message")
}
