package main

import "fmt"

func t12() {
	m := make(map[string]int)
	m["a"] = 1
	if v, ok := m["b"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("not exists")
	}
	//not exists
}

func main() {
	t12()
}
