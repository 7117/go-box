package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"name1" db:"name2"`
	Age  int64  `json:"age" bson:"age"`
}

func main() {
	var s Student
	s.Age = 10
	s.Name = "array"
	v := reflect.ValueOf(&s)
	fmt.Println("获取值:", v)
	// 类型
	t := v.Type()
	fmt.Println("获取类型:", t)
	// 获取字段
	f := t.Elem().Field(0)
	fmt.Println("获取json:", f.Tag.Get("json"))
	fmt.Println("获取db:", f.Tag.Get("db"))
}

//获取值: &{}
//获取类型: *main.Student
//获取字段: {Name  string json:"name1" db:"name2" 0 [0] false}
//获取json: name1
//获取db: name2
