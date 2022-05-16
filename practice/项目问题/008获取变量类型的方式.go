package main

import (
	"fmt"
	"reflect"
)

type Class struct {
	Id   int `json:"id"`
	Name int `json:"name"`
}

func main() {
	var class interface{}

	class = Class{
		1, 1,
	}

	使用switch断言的方式(class)
	使用反射的方式(class)
	使用格式化输出的方式(class)
}

func 使用switch断言的方式(class interface{}) {
	switch class.(type) {
	case int:
		fmt.Println("int")
	case Class:
		fmt.Println("Class")
	default:
		fmt.Println("default")
	}
}

func 使用反射的方式(class interface{}) {
	fmt.Println(reflect.TypeOf(class))
}

func 使用格式化输出的方式(class interface{}) {
	fmt.Println(fmt.Sprintf("%T", class))
}

//Class
//main.Class
//main.Class
