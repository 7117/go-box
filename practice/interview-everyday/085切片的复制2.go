package main

import "fmt"

func main() {
	var src, dst []int
	src = []int{1, 2, 3}
	dst = make([]int, len(src))
	n := copy(dst, src)
	fmt.Println(n, dst)
}

//3 [1 2 3]
