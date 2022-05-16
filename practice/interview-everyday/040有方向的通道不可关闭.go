package main

import "fmt"

func Stop(stop <-chan bool) {
	fmt.Println(stop)
	//本题目的目的：有方向的通道不可以被关闭
	//close(stop)
}

func main() {
	aaa := make(<-chan bool, 0)
	Stop(aaa)
}
