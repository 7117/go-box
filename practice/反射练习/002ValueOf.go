package main

import (
	"fmt"
	"reflect"
)

//反射获取interface值信息
func reflect_v1alue(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println(v)
	k := v.Kind()
	fmt.Println(k)
}

func main() {
	var x = 3.4
	reflect_v1alue(x)
}

//3.4
//float64
