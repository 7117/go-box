package main

import "fmt"

type User struct{}

// User1 创建了新类型User1
type User1 User

// User2 创建了User的别名User2
type User2 = User

func (i User1) m1() {
	fmt.Println("m1")
}
func (i User2) m2() {
	fmt.Println("m2")
}

func main() {
	var i1 User1
	var i2 User
	i1.m1()
	i2.m2()
}

//m1
//m2
