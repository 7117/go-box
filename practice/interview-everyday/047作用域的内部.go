package main

import "fmt"

func main() {
	x := 1
	fmt.Println(x)
	//出现在了作用域的内部
	{
		fmt.Println(x)
		i, x := 2, 2
		fmt.Println(i, x)
	}
	fmt.Println(x) // print ?
}

//1
//1
//2 2
//1
