package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	User := user{
		Id:   1,
		Name: "杉杉",
	}
	//第一种使用interface  强制转化
	var NewIF1 interface{}
	fmt.Println("1接口的类型", reflect.TypeOf(NewIF1))
	NewIF1 = User
	fmt.Println("1赋值的接口", NewIF1)
	fmt.Println("1赋值的类型", reflect.TypeOf(NewIF1))
	fmt.Printf("1使用interface的方法: %v\n", NewIF1.(user))
	//1接口的类型 <nil>
	//1赋值的接口 {1 杉杉}
	//1赋值的类型 main.user
	//1使用interface的方法: {1 杉杉}

	//第二种使用json
	var NewIF2 interface{}
	fmt.Println("2接口的类型", reflect.TypeOf(NewIF2))
	NewIF2 = User
	var newData user
	res, _ := json.Marshal(NewIF2)
	json.Unmarshal(res, &newData)
	fmt.Println("2赋值的接口", NewIF2)
	fmt.Println("2赋值的类型", reflect.TypeOf(NewIF2))
	fmt.Printf("2使用json的方法: %v\n", newData)
	//2接口的类型 <nil>
	//2赋值的接口 {1 杉杉}
	//2赋值的类型 main.user
	//2使用json的方法: {1 杉杉}

	//第三种断言  类似与第一种方法
	var NewIF3 interface{}
	NewIF3 = User
	p, ok := (NewIF3).(user)
	if ok {
		fmt.Println("3结构体:", p)
		fmt.Println("3id:", p.Id)
		fmt.Println("3name:", p.Name)
	} else {
		fmt.Println("3can not convert")
	}
	//3结构体: {1 杉杉}
	//3id: 1
	//3name: 杉杉
}
