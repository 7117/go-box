package main

import (
	"fmt"
	"runtime"
)

func main() {
	//饿死了
	runtime.GOMAXPROCS(1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	for {
		fmt.Println("BBBB")
	}

	fmt.Println("AAAA")
}
