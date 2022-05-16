package main

import (
	"fmt"
	"reflect"
)

func Foo(x interface{}) {
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.ValueOf(x))
	//这里的动态值为nil 但是动态类型是*int 所以是不相等
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
func main() {
	//定义了一个指针变量  这个指针变量为空git
	var x *int = nil
	Foo(x)
}

//*int
//<nil>
//non-empty interface
