package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println("go")
		}
	}()
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println("hello")
	}
}
