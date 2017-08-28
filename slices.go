package main

import (
	"fmt"
)

func slices() {
	//slice practice
	primes := [10]int{1, 2, 4, 5, 78, 1, 8, 9, 33}
	fmt.Println(primes)
	slice := primes[0:5]
	var v []int = primes[2:4]
	fmt.Println(slice)
	fmt.Println(v)

	//slice is reference of section of array
	p := primes[0:5]
	fmt.Println(p)
	p[0] = 100
	fmt.Println(p)
	fmt.Println(primes)

	//special slice
	fmt.Println("special slice")
	fmt.Println("0:5 --")
	s1 := primes[0:5]
	fmt.Println(s1)
	fmt.Println("0: --")
	s2 := primes[0:]
	fmt.Println(s2)
	fmt.Println(" :5 --")
	s3 := primes[:5]
	fmt.Println(s3)
	fmt.Println(" : --")
	s4 := primes[:]
	fmt.Println(s4)
}

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d , %v\n", len(s), cap(s), s)
}

func sliceLen() {
	primes := []int{1, 2, 4, 5, 78, 1, 8, 9, 33, 2}
	printSlice(primes)
	slice := primes[0:5]
	printSlice(slice)
	var v []int = primes[3:4]
	printSlice(v)
}

func nilSlice() {
	var v []int
	printSlice(v)
}

func makeSlice() {
	a := make([]int, 10)
	printSlice(a)
	b := make([]int, 3, 5)
	printSlice(b)
	c := b[:cap(b)]
	printSlice(c)
	d := b[1:]
	printSlice(d)
}

func slicesInSlices() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[1][1] = "X"
	board[2][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Println(board[i])
	}

}

func appending() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
}

func rangeF() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("index:%d, value:%d\n", i, v)
	}

	for i, _ := range pow {
		fmt.Printf("index:%d \n", i)
	}

	for _, i := range pow {
		fmt.Printf("index:%d \n", i)
	}
}

func mapF() {
	type Vertex struct {
		X float64
		Y float64
	}

	var m map[string]Vertex
	m = make(map[string]Vertex)

	m["hello"] = Vertex{
		7.231,
		-8.142,
	}

	fmt.Println(m["hello"])
	for a, b := range m {
		fmt.Println(a, b)
	}
}

func mapLiterals() {
	type Vertex struct {
		X, Y float64
	}
	var m = map[string]Vertex{
		"gong": {34.112, 45.876},
		"zhe":  {-34.112, -45.876},
	}

	for a, b := range m {
		fmt.Println(a, b)
	}
	fmt.Println(m)
}

func mapDelete() {
	var m map[string]int
	m = make(map[string]int)
	m["hello"] = 88
	fmt.Println(m)

	delete(m, "hello")
	fmt.Println(m)

	a, b := m["hello"]
	fmt.Println(a, b)
}
func main() {
	slices()
	/*
		sliceLen()
		nilSlice()
		makeSlice()
		slicesInSlices()
		appending()
		rangeF()
		mapF()
		mapLiterals()
		mapDelete()
	*/
}
