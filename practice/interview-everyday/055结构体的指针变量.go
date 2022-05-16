package main

import "fmt"

type WW struct {
	n int
}

func getT() WW {
	return WW{}
}

func main() {
	//直接返回的 WW{} 无法寻址，不可直接赋值。
	t := getT()
	fmt.Println(t)
	p := &t.n // <=> p = &(t.n)
	*p = 1
	fmt.Println(t.n)
}

//
//{0}
//1
