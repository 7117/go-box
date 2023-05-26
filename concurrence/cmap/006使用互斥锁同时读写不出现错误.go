package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	testLock()
}

func testLock() {
	//写入的时候使用读锁会导致：fatal error: concurrent map read and map write
	//读的时候使用读锁 写的时候使用互斥锁或者写锁
	loclVal := new(sync.RWMutex)

	m := make(map[int]int)
	go func() {
		for {
			loclVal.Lock()
			m[12] = 12
			m[100] = 100
			loclVal.Unlock()
		}
	}()
	go func() {
		for {
			loclVal.RLock()
			fmt.Println(m[12])
			fmt.Println(m[100])
			loclVal.RUnlock()
		}
	}()
	time.Sleep(4 * time.Second)
	fmt.Println("END")
}
