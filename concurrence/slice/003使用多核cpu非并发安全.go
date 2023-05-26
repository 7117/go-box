package main

import (
	"fmt"
	"runtime"
	"sync"
)

func concurrentAppendSliceNo() {
	sl := make([]int, 0)
	wg := sync.WaitGroup{}
	for index := 0; index < 100; index++ {
		wg.Add(1)
		go func(index int) {
			sl = append(sl, index)
			wg.Done()
		}(index)
	}
	wg.Wait()
	fmt.Println(sl)
	fmt.Printf("final len(sl)=%d cap(sl)=%d\n", len(sl), cap(sl))
}
func main() {
	runtime.GOMAXPROCS(8)
	concurrentAppendSliceNo()
	//final len(sl)=84 cap(sl)=128
}
