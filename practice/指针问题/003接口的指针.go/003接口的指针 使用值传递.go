package main

import "fmt"

type TInterfaced interface {
	M1()
	M2()
}

type T struct {
	Name string
}

func (t T) M1() {
}

func (t T) M2() {
}

func main() {
	var t1 = T{"t1"}
	t1.M1()
	t1.M2()

	//var t2 TInterfaced = &t1
	////&{t1}
	//fmt.Println(t2)
	var t2 TInterfaced = t1
	//{t1}
	fmt.Println(t2)
	t2.M1()
	t2.M2()
}
