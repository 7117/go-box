package main

import "fmt"

func main() {

	a := "a"
	defer func() {
		fmt.Println("内部", a)
	}()

	a = "c"
	fmt.Println("外部", a)
}

//外部 c
//内部 c
