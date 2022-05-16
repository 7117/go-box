package main

import "fmt"

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}
func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Println(elem)
	return s
}
func main() {
	s := NewSlice()
	fmt.Println("s值：", s)
	defer s.Add(1).Add(2)
	s.Add(3)
}

//1
//3
//2
//1.Add() 方法的返回值依然是指针类型 *Slice，所以可以循环调用方法 Add()；
//2.defer 函数的参数（包括接收者）是在 defer 语句出现的位置做计算的，而不是在函数执行的时候计算的，所以 s.Add(1) 会先于 s.Add(3) 执行。

//s值： []
