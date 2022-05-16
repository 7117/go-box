package main

import "fmt"

func f1() (r int) {
	r = 0
	//空参数才会影响到
	defer func() {
		r = r + 1
	}()
	return r
}
func f2() (a int) {
	a = 0
	//有形参的影响不到了
	defer func(a int) {
		a = a + 5
	}(a)

	return a
}
func f3() int {
	r := 0
	defer func() {
		r = r + 1
	}()
	return r
}
func main() {
	//1 0 0
	//返回参数没有定义变量+使用形参的时候 不会影响到返回值
	fmt.Println(f1(), f2(), f3())
}

//Go 中 defer 和 return 执行的先后顺序
//多个defer的执行顺序为“后进先出”
//defer、return、返回值三者的执行逻辑应该是：return最先执行，return负责将结果写入返回值中；接着defer开始执行一些收尾工作；最后函数携带当前返回值退出。

//如果函数的返回值是无名的（不带命名返回值），则go语言会在执行return的时候会执行一个类似创建一个临时变量作为保存return值的动作，
//而有名返回值的函数，由于返回值在函数定义的时候已经将该变量进行定义，在执行return的时候会先执行返回值保存操作，
//而后续的defer函数会改变这个返回值(虽然defer是在return之后执行的，但是由于使用的函数定义的变量，
//所以执行defer操作后对该变量的修改会影响到return的值
