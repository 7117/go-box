package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	//下面这一行的a已经是重新赋值后的a了
	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	//r是原来的  表示没有被重新赋值影响
	//as是新的   表示会被重新赋值
	//循环的是最远先的，原值会被影响
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}

//PS D:\goitem\go-collection\practice\interview-everyday> go run .\031.go
//r =  [1 2 3 4 5]
//a =  [1 12 13 4 5]
