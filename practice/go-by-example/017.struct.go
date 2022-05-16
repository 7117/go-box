package main

import "fmt"

type Persona struct {
	name string
	age  int
}

func main() {
	person := Persona{"张三", 12}

	fmt.Print(person)
}
