package main

import "fmt"

func main(){
    var i interface{}
    describe(i)    
 
    i=40
    describe(i)    
  
    i="ab"
    describe(i)    

}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
