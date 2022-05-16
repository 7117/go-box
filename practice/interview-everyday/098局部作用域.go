package main

import "fmt"

func main() {
	a := 1
	for i := 0; i < 5; i++ {
		a := a + 1
		fmt.Println(a)
		a = a * 2
		fmt.Println(a)
	}
	fmt.Println(a)
}

//1
