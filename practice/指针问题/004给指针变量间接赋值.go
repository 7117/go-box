package main

import "fmt"

func main() {
	var aa *int
	bb := 1

	aa = &bb
	fmt.Println(aa)
}

//PS D:\goitem\go-collection\practice\指针问题> go run .\004.go
//0xc00000e098 1
