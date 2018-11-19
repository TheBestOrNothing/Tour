package main

import (
    "fmt"
    "math/rand"
    "math"
)

func add(x , y int) int{
    return x+y
}

func swap(x, y string) (string, string){
    return x, y
}

func split(sum int) ( x, y int){
    x = sum
    y = sum +10
    return
}

var a, b, c int = 3,4,0

func main(){

    fmt.Println("This is main package", rand.Intn(10))
    fmt.Println(math.Pi)
    fmt.Println(add(12,33))
    x, y := swap("hello" , "world")
    fmt.Println(x,y)
    fmt.Println(split(1))
    var d, e = true, "NO" 
    fmt.Println(a,b,c,d,e)
    f:= false
    fmt.Println(a,b,c,d,e,f)
}

//In Go, a name is exported if it begins with a capital letter. 
//For example, Pizza is an exported name, as is Pi, which is exported from the math package. 
