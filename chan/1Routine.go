package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("go routinge message")
	fmt.Println("Main function message")
	time.Sleep(1 * time.Second)
}
