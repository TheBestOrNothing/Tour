package main

import "fmt"

func main() {
	var i interface{} = "hello world"
	s := i.(string)

	fmt.Println(s)

	fmt.Println(i.(string))

	f, ok := i.(string)
	fmt.Println(f, ok)

	g, ok := i.(int)
	fmt.Println(g, ok)

	fmt.Println(i.(int))

}
