package main 

import (
    "fmt"
    "math"
    "runtime"
    "time"
)

func for1(){
    sum := 0
    for i:=0;i<10;i++{
        sum += i
    }
    fmt.Println("Oupt for1",sum)
}

func for2(){
    sum := 1
    for sum <100 {
        sum += sum
    }
    fmt.Println("oupt for2",sum)
}

func for3(){
    sum := 1
    for {
        //sum += sum
    }
    fmt.Println("oupt for3",sum)
}

func ifstatement(x int) {
    if x<0 {
        fmt.Println(" X is less than zero")
    }else{
        fmt.Println(" X is not less than zero")
    }
}

func shortIf(x,n,lim float64) float64{
    if v := math.Pow(x,n); v<lim {
        return v
    }
    return lim
}

func caseStatement(){
    switch os:= runtime.GOOS; os{
        case "linux":
	    fmt.Println("Linux")
        case "windows":
	    fmt.Println("windows")
	default:
	    fmt.Println("The OS is",os)
    }
}

func switchTrue(){
    t:= time.Now()
    //switch true{
    switch{
       case t.Hour() < 11:
	   fmt.Println("Good Morning1")
       case t.Hour() < 12:
	   fmt.Println("Good Morning2")
       case t.Hour() < 17:
	   fmt.Println("Good afternoon")
       default:
	   fmt.Println("Good night")
    }
}

func deferStatement(){

   defer fmt.Println("Good")
   fmt.Println("Morning")
}

func stackDefer(){

   defer fmt.Println("stackDefer")
   for i:=0; i<10; i++ {
	defer fmt.Println(i)
   }
}

func main(){
   for1()
   for2()
   //for3()
   ifstatement(-1)
   ifstatement(1)
   fmt.Println(
      shortIf(2,3,20),
      shortIf(3,3,20),
   )
   caseStatement()
   switchTrue()
   deferStatement()
   stackDefer()
}
