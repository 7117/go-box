package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
}

func main() {
	u := &User{Id: 1001, Name: "xxx"}
	t := reflect.TypeOf(*u)
	v := reflect.ValueOf(*u)
	for k := 0; k < v.NumField(); k++ {
		fmt.Printf("%v -- %v \n", t.Field(k).Name, v.Field(k).Interface())
	}

	a := *u
	fmt.Println(a)
}
