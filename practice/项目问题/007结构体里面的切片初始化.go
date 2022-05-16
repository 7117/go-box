package main

import "fmt"

type Stu struct {
	Id   int   `json:"id"`
	Info []One `json:"info"`
}

type One struct {
	Name int `json:"name"`
	Addr int `json:"addr"`
}

func main() {
	var stu Stu
	info := make([]One, 5, 6)
	//此处make不可以进行赋值，没有这种形式
	info[0] = One{1, 2}
	stu.Info = info

	fmt.Println(stu)
}

//{0 [{1 2} {0 0} {0 0} {0 0} {0 0}]}
