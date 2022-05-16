package main

import (
	"fmt"
	"time"
)

func testssss(ch chan int) {
	fmt.Print("testssss1")
	ch <- 1
	fmt.Print("testssss2")
}

func main() {
	ch := make(chan int, 0)
	go testssss(ch)
	time.Sleep(3 * time.Second)
	fmt.Print("main end")
	<-ch
	time.Sleep(3 * time.Second)

}
