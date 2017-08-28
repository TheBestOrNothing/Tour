package main

import "fmt"

type Person struct{
    Name string
    Age  int
}

type Computer struct{
    CPU string
    MemSize int
}

func (p Person) String() string{
    return fmt.Sprintf("%v (%v year)", p.Name, p.Age)
}

func (c Computer) String() string{
    return fmt.Sprintf("CPU: %v Mem: %v ", c.CPU, c.MemSize)
}

func main() {
    p := Person{"gongzhe", 30}
    c := Computer{"2690V4",512 }
    fmt.Println(p,c)
}
