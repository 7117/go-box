package main

import (
	"encoding/json"
	"fmt"
)

type Stuss struct {
	Name string `json:"stuname"`
	Age  int
}

func main() {
	stu := Stuss{"a", 20}

	res, _ := json.Marshal(stu)

	fmt.Println(string(res))
}
