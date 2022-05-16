package main

import "fmt"

type N int

func (n N) test() {
	fmt.Println(n)
}

func main() {
	var n N = 10
	p := &n
	fmt.Println(p)

	n++
	f1 := n.test

	n++
	f2 := p.test

	n++
	fmt.Println(n)

	f1()
	f2()
}

//0xc0000aa058
//13
//11
//12
//参考答案及解析：
//13 11 12
//知识点：
//方法值，当指针值赋值给变量或者作为函数参数传递时，
//会立即计算并复制该方法执行所需的接收者对象，
//与其绑定，以便在稍后执行时，能隐式第传入接收者参数。
