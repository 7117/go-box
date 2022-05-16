package main

import "fmt"

type fooo struct {
	Val int
}

type bar struct {
	Val int
}

func main() {
	a := &fooo{Val: 5}
	b := &fooo{Val: 5}
	fmt.Println("bbbbb", b)
	c := fooo{Val: 5}
	d := bar{Val: 5}
	e := bar{Val: 5}
	f := bar{Val: 5}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(fooo(d))
	fmt.Println(e)
	fmt.Println(f)
	fmt.Print(a == b, c == fooo(d), e == f)
}

//bbbbb &{5}
//&{5}
//&{5}
//{5}
//{5}
//{5}
//{5}
//false true true
