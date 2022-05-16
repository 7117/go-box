package main

import "fmt"

func Readsss(ch chan int) {
	value := <-ch
	fmt.Println("value", value)
}

func main() {
	var chaaa chan int
	chaaa = make(chan int)
	Readsss(chaaa)
}
