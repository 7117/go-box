package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	var p Point         //正常
	fmt.Println(p)      //输出 {0 0}
	var nilPtr *Point   //报错
	fmt.Println(nilPtr) //<nil>
	var newPtr = new(Point)
	fmt.Println(newPtr) //&{0 0}

	//调用Abs()
	fmt.Println(p.Abs())      //正常
	fmt.Println(newPtr.Abs()) //正常
	fmt.Println(nilPtr.Abs()) //报错
}

//{0 0}
//<nil>
//&{0 0}
//
//0
//0
//panic: runtime error: invalid memory address or nil pointer dereference
