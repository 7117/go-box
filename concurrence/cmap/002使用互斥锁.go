package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//fatal error: all goroutines are asleep - deadlock!
	//此处是因为select{}   没有写入离开，就会导致死锁
	//所有的程序都处于休眠状态——死锁!
	count1 := 100
	count2 := 100
	loclVal := new(sync.Mutex)

	m := make(map[int]int)
	go func() {
		for {
			loclVal.Lock()
			m[12] = 12
			m[100] = 100
			fmt.Println("Write", count1)
			loclVal.Unlock()
			count1--
			if count1 <= 0 {
				break
			}
		}
		return
	}()
	go func() {
		for {
			loclVal.Lock()
			fmt.Println("Read", count2, m[12], m[100])
			loclVal.Unlock()
			count2--
			if count2 <= 0 {
				break
			}
		}
		return
	}()

	c := make(chan int)
	select {
	case <-c:
	default:
		time.Sleep(3 * time.Second)
	}
}
