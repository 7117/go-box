package main

import (
	"fmt"
	"runtime"
	"time"
)

func testdd() {
	defer fmt.Println("222")
	//终止这个协程
	runtime.Goexit()
	fmt.Println("333")
}

func main() {
	go func() {
		fmt.Println("111")
		testdd()
		fmt.Println("444")
	}()

	time.Sleep(5 * time.Second)

}
