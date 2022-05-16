package main

import "fmt"

func main() {

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		//相当于新建了一个变量
		value := val
		m[key] = &value
	}

	fmt.Println(m)

	for k, value := range m {
		fmt.Println(k, "===>", *value)
	}
}
