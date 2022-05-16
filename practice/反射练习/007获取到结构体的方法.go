package main

import (
	"fmt"
	"reflect"
)

// User4 定义结构体
type User4 struct {
	Id   int
	Name string
	Age  int
}

func (u User4) Hello(name string) {
	fmt.Println("Hello：", name)
}

func main() {
	u := User4{1, "5lmh.com", 20}

	v := reflect.ValueOf(u)
	fmt.Println("获取到结构体的值:", v)

	m := v.MethodByName("Hello")
	fmt.Println("获取到结构体方法:", m)

	// 构建一些参数
	args := []reflect.Value{reflect.ValueOf("6666")}
	fmt.Println(args)
	// 没参数的情况下：var args2 []reflect.Value
	// 调用方法，需要传入方法的参数
	m.Call(args)
}

//{1 5lmh.com 20}
//0xff8fe0
//[6666]
//Hello： 6666
