package main

import "fmt"

func t1() {
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:3]

	fmt.Println(a, b, c)
	fmt.Println(cap(s))
	fmt.Println(cap(c))
	//[] [1 2] [2]
	//3
	//2
}

func main() {
	t1()
}
