package main

import (
	"fmt"
	"reflect"
)

// User1 定义结构体
type User1 struct {
	Id   int
	Name string
	Age  int
}

// Boy 匿名字段
type Boy struct {
	User1
	Addr string
}

func main() {
	m := Boy{User1{1, "zs", 20}, "bj"}
	t := reflect.TypeOf(m)
	fmt.Println(t)
	// Anonymous：匿名
	fmt.Printf("%#v\n", t.Field(0))

	// 值信息
	fmt.Printf("%#v\n", reflect.ValueOf(m).Field(0))
}

//main.Boy
//reflect.StructField{Name:"User1", PkgPath:"", Type:(*reflect.rtype)(0xdce280), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:true}
//main.User1{Id:1, Name:"zs", Age:20}
