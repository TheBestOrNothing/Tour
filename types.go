package main

import(
    "fmt"
)

func pointer(){
    var p *int
    var q *int
    i := 10
    p = &i

    fmt.Println("Address:", p)
    fmt.Println("The value:", *p)

    *p = 12
    fmt.Println("The value:", *p)
    fmt.Println("The value:", q)
}

func structs(){
    //Define a struct
    type Vertex struct{
    	X int
	Y int
    }

    fmt.Println(Vertex{1,2})

    //To use structs
    V := Vertex{12,34}
    fmt.Println(V)
    V.X = 44
    fmt.Println(V)

    //Pointers to structs
    V1 := Vertex{3,5}
    fmt.Println(V1)
    var P2V *Vertex
    P2V = &V1
    P2V.Y = 9
    fmt.Println(V1)
}

func structLiterals(){

    type Vertex struct{
        X int
	Y int
    }

    v1 := Vertex{1,2}
    v2 := Vertex{X:1}
    v3 := Vertex{Y:1}
    p  := &Vertex{3,1}
    fmt.Println(v1,v2,v3,p)
}

func arrays(){
    var arr [2]string
    arr[0] = "hello"
    arr[1] = "world"
    fmt.Println(arr);
    
    primes := [10]int{1,2,2,3,4}
    fmt.Println(primes);
    
}

func slices(){
    //slice practice
    primes :=  [10]int{1,2,4,5,78,1,8,9,33}
    fmt.Println(primes)
    slice  :=  primes[0:5]
    var v []int =  primes[3:4]
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

func printSlice(s []int){
    fmt.Printf("len=%d, cap=%d , %v\n", len(s),cap(s),s)
}

func sliceLen(){
    primes :=  []int{1,2,4,5,78,1,8,9,33,2}
    printSlice(primes)
    slice  :=  primes[0:5]
    printSlice(slice)
    var v []int =  primes[3:4]
    printSlice(v)
}

func nilSlice(){
    var v []int
    printSlice(v)
}

func makeSlice(){
    a := make([]int, 10)
    printSlice(a)
    b := make([]int, 3, 5)
    printSlice(b)
    c := b[:cap(b)]
    printSlice(c)
    d := b[1:]
    printSlice(d)
}

func slicesInSlices(){
    board := [][]string{
        []string{"_","_","_"},
        []string{"_","_","_"},
        []string{"_","_","_"},
    }

    board[0][0] = "X"
    board[1][1] = "X"
    board[2][2] = "X"
   
    for i:=0;i<len(board);i++{
        fmt.Println(board[i])
    }

}

func appending(){
    var s [] int
    printSlice(s)

    s = append(s,0)
    printSlice(s)
    
    s = append(s,1)
    printSlice(s)

    s = append(s,2,3,4)
    printSlice(s)
}

func main(){
   pointer()
   structs()
   structLiterals()
   arrays()
   slices()
   sliceLen()
   nilSlice()
   makeSlice()
   slicesInSlices()
   appending()
}
