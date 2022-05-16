package main

import "fmt"

func main() {
	lists := new([]int)
	lists = append(lists, 1)
	fmt.Println(lists)
}
