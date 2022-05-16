package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("我是第一个")
		fmt.Print(recover())
	}()
	defer func() {
		fmt.Println("我是第二个")
		defer fmt.Print(recover())
		fmt.Println()
		panic(1)
	}()
	//此处没有在函数中 所以是错误的
	//defer recover()
	fmt.Println("我是第三个")
	panic(2)
}

//我是第三个
//我是第二个
//
//2我是第一个
//1
