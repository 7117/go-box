package main

import "fmt"

type Person struct {
	age int
}

func main() {
	person := &Person{28}

	// 1.直接给赋值了  就是28
	defer fmt.Println(1, person.age)

	// 2.实参是已经给复制了   到达形参后   就没法影响了  是建立一个临时变量
	//他是函数内部哈
	defer func(p *Person) {
		fmt.Println(2, p.age)
	}(person)

	// 3.
	defer func() {
		fmt.Println(3, person.age)
	}()

	person = &Person{29}
}

//3 29
//2 28
//1 28
