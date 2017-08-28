package main

import (
    "fmt"
)


type Vertex struct{
    X,Y int
}

type FLOAT  float64

func add1(x,y int) int{
    return x+y
}

func add2(v Vertex) int{
    return v.X + v.Y
}

func (v Vertex) add3() int {
    return v.X + v.Y
}

func add4() FLOAT{
    var f1  FLOAT
    var f2  FLOAT
    f1 = 1.2
    f2 = 4.1
    return  f1 + f2
}

func (v *Vertex) add5(f int) Vertex{
    v.X = f
    v.Y = f*2
    return *v
}

func main(){
    v1 := Vertex{X:1, Y:2}
    a1 := add1(v1.X, v1.Y)
    fmt.Println(a1)

    a2 := add2(v1)
    fmt.Println(a2)

    fmt.Println(v1.add3())
    fmt.Println(add4())
    fmt.Println(v1)
    fmt.Println(v1.add5(5))
    fmt.Println(v1)
}

