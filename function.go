package main

import(
    "fmt"
)

func compute(fn func(x,y int) int) int{
    return fn(3,4)
}

func funcValue(){
    add := func(x,y int) int{
	return x+y
    }
    fmt.Println(add(3,4))

    fmt.Println(compute(add))

}

func adder() func(int) int{
    sum := 0
    return func(x int) int{
	sum += x
	return sum
    }
}

func sum(){
    pos, neg := adder(),adder()
    for i:=0;i<10;i++ {
	fmt.Println(
	    pos(i),
	    neg(-i),
	)
    }
}

func main(){
    funcValue()
    sum()
}
