package main

import "fmt"

func main() {
	x := []int{0, 1, 2}
	y := [3]*int{}
	for i, v := range x {
		defer func() {
			fmt.Println(v)
		}()
		y[i] = &v
	}
	fmt.Println(*y[0])
	fmt.Println(*y[1])
	fmt.Println(*y[2])
}

//2
//2
//2
//2
