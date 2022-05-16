package main

import "fmt"

type Persond struct {
	name string
	age  int
}

type Studenta struct {
	Persond
	num string
}

func main() {
	person := Persond{"张三", 12}
	fmt.Print(person)

	student := Studenta{Persond{"账务", 11}, "student"}
	fmt.Print(student)
	fmt.Print(student.num)
	fmt.Print(student.Persond.age)
}
