package main

import "fmt"

type TimesMatcherE struct {
	base int
}

func NewTimesMatcherE(base int) TimesMatcherE {
	return TimesMatcherE{base: base}
}

func main() {
	p := NewTimesMatcherE(3)
	fmt.Println(p)
}

//{3}
//Go语言的内存回收机制规定，只要有一个指针指向引用一个变量，那么这个变量就不会被释放（内存逃逸），因此在 Go 语言中返回函数参数或临时变量是安全的
