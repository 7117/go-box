package main

import "fmt"

func main() {
	////方式一：短声明
	//a := []struct {
	//	id   int64
	//	name string
	//}{
	//	{
	//		id:   1,
	//		name: "aaaa",
	//	},
	//	{
	//		id:   2,
	//		name: "bbbb",
	//	},
	//}
	//fmt.Println(a)
	//PS D:\goitem\go-collection\view\其他练习> go run .\002[]struct.go
	//[{1 aaaa} {2 bbbb}]

	////方式二：定义变量
	//type c struct {
	//	id   int64
	//	name string
	//}
	//
	//var b []c
	//
	//b = make([]c, 2, 100)
	//
	//b[0].id = 1
	//b[0].name = "a"
	//fmt.Println(b)
	////PS D:\goitem\go-collection\view\其他练习> go run .\002[]struct.go
	////[{1 a} {0 }]

	fmt.Println("end")
}
