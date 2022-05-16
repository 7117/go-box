package main

import "fmt"

func main() {
	var a []int
	a = make([]int, 1, 111)
	a[0] = 1
	a = append(a, []int{1, 2, 3, 34}...)
	fmt.Println(a)
}
