package main

import "fmt"
import "time"

func Addaa(x, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	for i := 0; i < 10; i++ {
		go Addaa(i, i)
	}

	time.Sleep(10)
}
