package main

import (
	"fmt"
	"sync"
)

func cChan() {
	sl := make(chan int, 100)
	wg := sync.WaitGroup{}
	for index := 0; index < 100; index++ {
		wg.Add(1)
		go func(index int) {
			sl <- index
			wg.Done()
		}(index)
	}
	wg.Wait()
	//final len(sl)=100 cap(sl)=100
	fmt.Printf("final len(sl)=%d cap(sl)=%d\n", len(sl), cap(sl))
}
func main() {
	cChan()
}
