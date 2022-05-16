package main

import "fmt"

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type C struct {
	i int
}

func (c C) ShowA() int {
	return c.i + 10
}

func (c C) ShowB() int {
	return c.i + 20
}

//这边就是多态 接口  可以  被同样的结构体  进行赋值
func main() {
	c := C{3}
	var a A
	a = c
	var b B
	b = c
	//13
	//23
	fmt.Println(a.ShowA())
	fmt.Println(b.ShowB())
}
