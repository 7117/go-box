package main

import (
	"encoding/json"
	"fmt"
)

type Useraa struct {
	Id   int
	Name string
	Age  interface{}
}

func main() {
	u := &Useraa{Id: 1001, Name: "xxx", Age: "11111111"}
	sss := u.Age
	fmt.Println(sss)
	marshal, _ := json.Marshal(*u)
	fmt.Println(string(marshal))
}
