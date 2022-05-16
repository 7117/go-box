package main

import "fmt"

type T struct {
	Name string
}

func (t T) M1() {
	t.Name = "name1"
}

// M2 *T就是指针变量  就是t   t是一个指针变量  存放的是地址  地址的Name字段是一个字符串  表明修改的是地址上的字段
func (t *T) M2() {
	//2分析指针变量的类型
	//输出类型 *main.T---*表示是一个指针变量，后面的T表示的是指针变量所指向的数据的类型  所以*T是和变量对等的
	//输出值的 &{t1}-----是一个地址，传递的是t1的地址
	fmt.Println(fmt.Sprintf("%T", t))
	fmt.Println(t)
	//就是地址的name字段
	t.Name = "name2"
}

func main() {
	t1 := T{"t1"}

	//因为形参会建立一个临时变量   函数里面的t与外面的t不一致了，所以说不会进行更改了！！！！！！
	fmt.Println("M1调用前：", t1.Name)
	t1.M1()
	fmt.Println("M1调用后：", t1.Name)

	fmt.Println("M2调用前：", t1.Name)
	t1.M2()
	fmt.Println("M2调用后：", t1.Name)
}

//M1调用前： t1
//M1调用后： t1
//M2调用前： t1
//M2调用后： name2
