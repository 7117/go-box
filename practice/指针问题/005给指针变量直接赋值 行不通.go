package main

import "fmt"

func main() {
	var aa *int

	*aa = 1

	fmt.Println(aa)
}

//panic: runtime error: invalid memory address or nil pointer dereference
//[signal 0xc0000005 code=0x1 addr=0x0 pc=0x107c516]
//
//goroutine 1 [running]:
//main.main()
//D:/goitem/go-collection/practice/指针问题/005给指针变量直接赋值.go:8 +0x16
//exit status 2
