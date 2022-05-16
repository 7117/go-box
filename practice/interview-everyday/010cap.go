package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	//长度：j-i，容量：k-i
	t := a[3:4:4]
	fmt.Println(t, t[0], cap(t))
}
