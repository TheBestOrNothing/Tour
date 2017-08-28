package main

import "fmt"

func sum(arr []int, c chan int) {
	sum := <-c
	fmt.Printf("In sum ")

	fmt.Println(len(arr))
	for _, v := range arr {
		sum += v
	}
	fmt.Println(sum)
}

func main() {
	//array := []int{1, 2, 4, 5, 6, -8, 0, 3}
	c := make(chan int, 1)
	c <- 6
	val := <-c
	fmt.Println(val)

	fmt.Println("Done")
}
