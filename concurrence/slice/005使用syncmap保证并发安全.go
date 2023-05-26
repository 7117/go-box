package main

import (
	"fmt"
	"sync"
)

func concurrentSyncMap() {
	var sl sync.Map

	wg := sync.WaitGroup{}
	for index := 0; index < 100; index++ {
		wg.Add(1)
		go func(index int) {
			sl.Store(index, index)
			wg.Done()
		}(index)
	}
	wg.Wait()
	fmt.Printf("final len(sl)=%d cap(sl)=%d\n", GetSyncMapCount(&sl))
}
func main() {
	concurrentSyncMap()
}

func GetSyncMapCount(scene *sync.Map) int64 {
	count := 0
	scene.Range(func(k, v interface{}) bool {
		count++
		return true
	})
	return int64(count)
}
