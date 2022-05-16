package main

import "fmt"

func main() {
	x := interface{}(nil)
	y := (*int)(nil)
	a := y == x
	//y <nil>
	fmt.Println("y", y)
	b := y == nil
	//x <nil>  此时x的动态类型为nil
	fmt.Println("x", x)
	_, c := x.(interface{})
	println(a, b, c)
}

//false true false
//类型断言语法：i.(Type)，其中 i 是接口，Type 是类型或接口。
//编译时会自动检测 i 的动态类型与 Type 是否一致。
//但是，如果动态类型不存在，则断言总是失败
