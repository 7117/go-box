package main

import (
	"fmt"
	"time"
)

func testr(ch chan int) {
	fmt.Print("testr1")
	ch <- 1
	fmt.Print("testr2")
}

func main() {
	ch := make(chan int, 1)
	go testr(ch)
	time.Sleep(1 * time.Second)
	fmt.Print("mainend")
	<-ch
	time.Sleep(1 * time.Second)

}
