package main

import "fmt"

type SliceAA []int

func NewSliceAA() SliceAA {
	return make(SliceAA, 0)
}
func (s *SliceAA) Add(elem int) *SliceAA {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}
func main() {
	s := NewSliceAA()
	defer func() {
		s.Add(1).Add(2)
	}()
	s.Add(3)
}

//312。对比昨天的第二题，本题的 s.Add(1).Add(2) 作为一个整体包在一个匿名函数中，会延迟执行。
