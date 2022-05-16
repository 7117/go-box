package main

import "fmt"

func main() {
	a := [3]int{0, 1, 2}
	s := a[1:2]
	fmt.Println("s1:", s)

	s[0] = 11
	s = append(s, 12)
	s = append(s, 13)
	s[0] = 21

	fmt.Println(a)
	fmt.Println(s)
}

//s1: [1]
//[0 11 12]
//[21 12 13]
