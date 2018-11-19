package main

import(
    "fmt"
)

type Vertex struct{
    x,y Float64 
}
type Float64 float64

type plus interface{
    change(Float64)
}

func (i *Float64) change(plus Float64) {
    if i == nil {
	fmt.Println("<nil> is <nil>")
	return
    }
    *i =  plus
}

func (v *Vertex) change(plus Float64) {
    v.x = plus
    v.y = plus
}

func main(){
    var p plus
    var a Float64
    var c *Float64
    a=2.01
    b:=Vertex{1,2}
    
    fmt.Println(a)
    p = &a
    p.change(10)
    fmt.Println(a)
    
    fmt.Println(b)
    p = &b
    p.change(10)
    fmt.Println(b)

    fmt.Println(c)
    p = c
    p.change(10)
    fmt.Println(c)
}

