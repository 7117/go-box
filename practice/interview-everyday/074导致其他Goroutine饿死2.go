package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	//没饿死
	runtime.GOMAXPROCS(1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		os.Exit(0)
	}()

	select {}
}
