package main

import "fmt"

func main() {

	type a struct {
		aa  int
		aaa string
	}

	type b struct {
		bb   int
		bbb  string
		bbbb a
	}

	//方法一：声明类型的进行赋值
	bbbbb := b{bb: 1, bbb: "bbb", bbbb: a{aa: 1, aaa: "aaa"}}
	fmt.Println(bbbbb)

	//	方法二：进行直接等于号赋值
	var c b
	c.bb = 1
	c.bbbb.aa = 1
	fmt.Println(c)

}
