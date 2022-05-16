package main

import "fmt"

var i int64 = 1

func main() {
	for {
		i++
		fmt.Println(i)
		return
	}
}
