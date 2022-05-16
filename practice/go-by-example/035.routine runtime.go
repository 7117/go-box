package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		for i := 1; i < 5; i++ {
			if i == 2 {
				runtime.Gosched()
			}
			fmt.Println(i)
		}
	}()

	go func() {
		for i := 6; i <= 8; i++ {
			fmt.Println(i)
		}
	}()

	time.Sleep(2 * time.Second)
}
