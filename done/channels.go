package main

import "fmt"

func sum(arr []int, c chan int) {
	sum := 0

	fmt.Println(len(arr))
	for _, v := range arr {
		sum += v
	}
	c <- sum
}

func buffered() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func buffered1() {
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println(<-ch)
	ch <- 2
	fmt.Println(<-ch)
}

func main() {
	array := []int{1, 2, 4, 5, 6, -8, 0, 3}
	c := make(chan int, 1)
	go sum(array[len(array)/2:], c)
	go sum(array[:len(array)/2], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
	buffered()
	buffered1()
}
