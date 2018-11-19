package main


import(
    "fmt"
)


type Vertex struct{
    x,y int
}

func (v *Vertex) change1(){
    v.x = 10
    v.y = 20
}

func (v Vertex) change2(){
    v.x = 30
    v.y = 40
}

func main(){
    v1 := Vertex{1,2}
    fmt.Println(v1)
    v1.change1()
    fmt.Println(v1)

    v2 := Vertex{3,4}
    fmt.Println(v2)
    v2.change2()
    fmt.Println(v2)
}
