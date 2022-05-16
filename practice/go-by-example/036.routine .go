package main

import (
	"fmt"
	"time"
)

var ch chan int

func main() {

	ch = make(chan int)

	go func() {
		for i := 1; i < 4; i++ {
			if i == 2 {
				<-ch
			}
			fmt.Println(i)
		}
	}()

	go func() {
		for i := 6; i <= 8; i++ {
			fmt.Println(i)
		}
		ch <- 1
	}()

	time.Sleep(2 * time.Second)
}
