package main

import (
	"fmt"
	"reflect"
)

// ReflectSetValue 反射修改值
func ReflectSetValue(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println("类型:", fmt.Sprintf("%T", v))
	fmt.Println("类型:", reflect.TypeOf(v))
	fmt.Println("类型:", v.Kind())

	switch v.Kind() {
	case reflect.Float64:
		// 反射修改值
		v.SetFloat(6.9)
		fmt.Println("值是什么值:", v.Float())
	case reflect.Ptr:
		// Elem()获取地址指向的值
		v.Elem().SetFloat(7.9)
		// 值
		fmt.Println("获取值v.Elem().Float():", v.Elem().Float())
		// 地址
		fmt.Println("获取指针v.Pointer()：", v.Pointer())
	}
}

func main() {
	var x = 3.4
	// 反射认为下面是指针类型，不是float类型
	ReflectSetValue(&x)
	fmt.Println("函数修好后的值:", x)
}

//类型: reflect.Value
//类型: reflect.Value
//类型: ptr
//获取值v.Elem().Float(): 7.9
//获取指针v.Pointer()： 824633761944
//函数修好后的值: 7.9
