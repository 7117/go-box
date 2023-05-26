package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

func concurrentAppendSliceNotForceIndexb() {
	sl := make([]int, 0)
	wg := sync.WaitGroup{}
	for index := 0; index < 100; index++ {
		wg.Add(1)
		go func(index int) {
			lock.Lock()
			sl = append(sl, index)
			lock.Unlock()
			wg.Done()
		}(index)
	}
	wg.Wait()
	fmt.Println(sl)
	fmt.Printf("final len(sl)=%d cap(sl)=%d\n", len(sl), cap(sl))
}
func main() {
	concurrentAppendSliceNotForceIndexb()
	//使用的001案例，加上锁之后，不在出现了数量不对了，都是100个
}
