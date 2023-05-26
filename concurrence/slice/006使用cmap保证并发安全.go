package main

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map/v2"
	"sync"
)

func concurrentcmap() {
	sl := cmap.New[int64]()

	wg := sync.WaitGroup{}
	for index := 0; index < 100; index++ {
		wg.Add(1)
		go func(index int) {
			sl.Set(string(index), int64(index))
			wg.Done()
		}(index)
	}
	wg.Wait()
	fmt.Println(sl)
	fmt.Printf("final len(sl)=%d cap(sl)=%d\n", sl.Count())
}
func main() {
	concurrentcmap()

}
