package main

import "fmt"

type Math struct {
	x, y int
}

var m = map[string]*Math{
	"foo": {2, 3},
}

func main() {
	m["foo"].x = 4
	fmt.Println(m["foo"].x)
	fmt.Printf("%#v", m["foo"]) // %#v 格式化输出详细信息
	fmt.Println()
	fmt.Printf("%+v", m["foo"]) // %#v 格式化输出详细信息
}

//4
//&main.Math{x:4, y:3}
//&{x:4 y:3}
