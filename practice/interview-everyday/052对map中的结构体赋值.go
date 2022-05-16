package main

import "fmt"

type R struct {
	n int
}

func main() {
	m := make(map[int]R)
	//map[key]struct 中 struct 是不可寻址的，所以无法直接赋值。
	//应该先给结构体进行赋值再进行其他的
	ta := R{1}
	m[0] = ta
	fmt.Println(m[0].n)
}
