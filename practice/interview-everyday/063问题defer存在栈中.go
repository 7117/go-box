package main

import "fmt"

func F(n int) func() int {
	return func() int {
		n++
		return n
	}
}

func main() {
	f := F(5)
	defer func() {
		fmt.Println("eeeee", f())
	}()
	defer fmt.Println("wwwww", f())
	i := f()
	fmt.Println("qqqqq", i)
}

//qqqqq 7
//wwwww 6
//eeeee 8
//匿名函数、defer()。defer() 后面的函数如果带参数，会优先计算参数，并将结果存储在栈中，到真正执行 defer() 的时候取出。
