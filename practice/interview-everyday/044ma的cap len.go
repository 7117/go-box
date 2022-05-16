package main

import "fmt"

func main() {
	m := make(map[string]int, 2)
	m["a"] = 1
	m["b"] = 1
	//使用cap对于map是错误的
	//fmt.Println(cap(m))

	//使用len获取map的容量
	//2
	fmt.Println(len(m))
}
