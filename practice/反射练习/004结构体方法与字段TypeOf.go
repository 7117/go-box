package main

import (
	"fmt"
	"reflect"
)

// User 定义结构体
type User struct {
	Id   int
	Name string
	Age  int
}

// Hello 绑方法
func (u User) Hello() {
	fmt.Println("hah")
}

// Ponzi 传入interface{}
func Ponzi(o interface{}) {
	t := reflect.TypeOf(o)

	fmt.Println("=================结构体信息====================")
	fmt.Println("o的类型是：", t)
	fmt.Println("字符串表示：", t.Name())

	fmt.Println("=================获取值====================")
	// 获取值
	v := reflect.ValueOf(o)
	fmt.Println(v)
	// 获取结构体字段个数：t.NumField()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Println(fmt.Sprintf("%s : %v : %v", f.Name, f.Type, val))
		// 获取字段的值信息 Interface()：获取字段对应的值
	}
	fmt.Println("=================方法====================")

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
		fmt.Println(m.Type)
	}

}

func main() {
	u := User{1, "zs", 20}
	Ponzi(u)
}

//=================结构体信息==============
//o的类型是： main.User
//字符串表示： User
//=================获取值==================
//{1 zs 20}
//Id : int : 1
//Name : string : zs
//Age : int : 20
//=================方法====================
//Hello
//func(main.User)
