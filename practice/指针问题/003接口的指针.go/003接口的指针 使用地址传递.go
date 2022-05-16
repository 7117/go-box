package main

import (
	"fmt"
	"reflect"
)

type TAInterfaced interface {
	M1()
	M2()
}

type TA struct {
	Name string
}

func (t TA) M1() {
	fmt.Println(reflect.ValueOf(t))
}

func (t *TA) M2() {
}

func main() {
	var t1 = TA{"t1"}

	var t2 TAInterfaced = &t1
	fmt.Println(t2)
	t2.M1()
	t2.M2()
}
