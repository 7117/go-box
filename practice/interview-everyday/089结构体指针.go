package main

import "fmt"

type P *int
type Q *int

func main() {
	var p P = new(int)
	fmt.Println("p,*p:", p, *p)
	III := *p + 8
	*p = *p + 8
	fmt.Println("III,*p:", III, *p)
	var x *int = p
	fmt.Println("x:", x)
	var q Q = x
	*q++
	fmt.Println("*p, *q:", *p, *q)
}

//p,*p: 0xc00000a098 0
//III,*p: 8 8
//x: 0xc00000a098
//*p, *q: 9 9
