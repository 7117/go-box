package main

import "fmt"

func testdata() []func() {
	var funs []func()
	// 0 1
	for i := 0; i < 2; i++ {
		funs = append(funs, func() {
			fmt.Println("子函数1：", &i, i)
		})
		fmt.Println("子函数2：", funs)
	}
	return funs
}

func main() {
	funs := testdata()
	fmt.Println("主函数的:", funs)
	for _, f := range funs {
		f()
	}
}

//子函数2： [0x119c6e0]
//子函数2： [0x119c6e0 0x119c6e0]
//主函数的: [0x119c6e0 0x119c6e0]
//子函数1： 0xc00000a098 2
//子函数1： 0xc00000a098 2
//知识点：
//闭包延迟求值,for循环局部变量i，匿名函数每一次使用的都是同一个变量。（说明：i 的地址，输出可能与上面的不一样）。
