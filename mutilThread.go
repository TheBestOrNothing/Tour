package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	defer fmt.Printf("main defer \n")
	go test1()
	time.Sleep(3000000)

}
func test() {
	num := 0
	//mtx := &sync.Mutex{}
	var wg sync.WaitGroup

	for idx := 1; idx < 50; idx++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			//mtx.Lock()
			num++
			fmt.Printf("num: %d, idx: %d\n", num, idx)
			//mtx.Unlock()
		}(idx)
	}
	wg.Wait()
	fmt.Printf("lastline num is %d\n", num)

}

func test1() {

	defer fmt.Printf("test1 defer \n")
	num := 0
	mtx := &sync.Mutex{}

	for idx := 1; idx < 50; idx++ {
		go func(idx int) {
			mtx.Lock()
			num++
			//fmt.Printf("num: %d, idx: %d\n", num, idx)
			mtx.Unlock()
			if num > 20 {
				fmt.Printf("num is  %d\n", num)
			}
		}(idx)
	}
	fmt.Printf("lastline num is %d\n", num)

}
