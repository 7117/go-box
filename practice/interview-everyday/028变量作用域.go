package main

import "fmt"

var p *int

func foo() (*int, error) {
	var i = 5
	return &i, nil
}

func main() {
	p, _ := foo()

	//这个p存储的是地址
	fmt.Println(p)
	//这个*表示的是找到p存储的地址所指向的变量
	//*另外的表示就是表示是一个指针变量
	fmt.Println(*p)
}

//PS D:\goitem\go-collection\practice\interview-everyday> go run .\028变量作用域.go
//0xc00000e098
//5
