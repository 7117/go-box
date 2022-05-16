package main

import (
	"fmt"
	"reflect"
)

// User3 定义结构体
type User3 struct {
	Id   int
	Name string
	Age  int
}

// SetValue 修改结构体值
func SetValue(o interface{}) {
	v := reflect.ValueOf(o)
	fmt.Println("通过valueOf进行获取到元素", v)
	v = v.Elem()
	fmt.Println("获取指针指向的元素", v)
	// 取字段
	f := v.FieldByName("Name")
	fmt.Println("进行FieldByName获取到字段值", f)
	if f.Kind() == reflect.String {
		f.SetString("Cutting")
	}
}

func main() {
	u := User3{1, "5lmh.com", 20}
	SetValue(&u)
	fmt.Println("修改结构体值", u)
}

//
//通过valueOf进行获取到元素 &{1 5lmh.com 20}
//获取指针指向的元素 {1 5lmh.com 20}
//进行FieldByName获取到字段值 5lmh.com
//修改结构体值 {1 Cutting 20}
