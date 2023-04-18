package main

import (
	"fmt"
	"time"
)

func main() {

	b0 := time.Now()
	fmt.Println("b0:", b0)
	b00 := b0.Format("2006-01-02 15:04:05")
	fmt.Println("b00:", b00)

	b1 := time.Now().UTC()
	fmt.Println("b1:", b1)
	b11 := b1.Format("2006-01-02 15:04:05")
	fmt.Println("b11:", b11)

	b3 := time.Now()
	fmt.Println("b3:", b3)
	fmt.Println("b3:", b3.Local())
	fmt.Println("b3:", b3.Location())
	b33 := b3.Format("2006-01-02 15:04:05")
	fmt.Println("b33:", b33)
}
