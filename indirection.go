package main

import(
    "fmt"
)

type Vertex struct{
   x,y int
}

func (v Vertex) add1() int {
    v.x = 9
    v.y = 6
    return v.x + v.y
}

func (v *Vertex) add2() int {
    v.x = 9
    v.y = 6
    return v.x + v.y
}

func main(){
    v1 := Vertex{1,1}
    v2 := &Vertex{2,2}

    fmt.Println(v1)
    fmt.Println(v2)
    fmt.Println(v1.add2())
    fmt.Println(v2.add2())
    fmt.Println(v1)
    fmt.Println(v2)
}
