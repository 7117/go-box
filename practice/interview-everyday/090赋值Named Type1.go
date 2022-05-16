package main

import "fmt"

type TTT int

func FFFF(t TTT) {}

func main() {
	var qq int
	//此处会报错  无法编译
	//FFFF(qq)
	fmt.Println(qq)
}
