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

func main() {
	array := []int{1, 2, 4, 5, 6, -8, 0, 3}
	c := make(chan int, 1)
	fmt.Printf("Start sum .......")
	sum(array[len(array)/2:], c)
	fmt.Printf("half sum .......")
	sum(array[:len(array)/2], c)
	fmt.Printf("anohter half sum .......")

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
