package main

import "fmt"

func f11() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f22() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()

	return t
}

func f33() (a int) {
	a = 1
	defer func(a int) {
		a = a + 5
	}(a)

	return a
}

func f4() (t int) {
	t = 5
	defer func() {
		t = t + 5
	}()
	return t
}
func main() {
	//1 5 1 10
	fmt.Println(f11(), f22(), f33(), f4())
}

//Go 中 defer 和 return 执行的先后顺序
//多个defer的执行顺序为“后进先出”
//
//defer、return、返回值三者的执行逻辑应该是：
//return最先执行，return负责将结果写入返回值中；接着defer开始执行一些收尾工作；最后函数携带当前返回值退出。

//如果函数的返回值是无名的（不带命名返回值），则go语言会在执行return的时候会执行一个类似创建一个临时变量作为保存return值的动作，
//而有名返回值的函数，由于返回值在函数定义的时候已经将该变量进行定义，在执行return的时候会先执行返回值保存操作，
//而后续的defer函数会改变这个返回值(虽然defer是在return之后执行的，但是由于使用的函数定义的变量，
//所以执行defer操作后对该变量的修改会影响到return的值
