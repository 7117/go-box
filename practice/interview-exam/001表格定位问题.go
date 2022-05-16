package main

import "fmt"

func main() {
	a := [][]int{{1, 2}, {3, 4}}
	fmt.Println(a)

	b := map[string]int{"a": 1, "b": 2}
	fmt.Println(b)

	var e map[string]int
	e = make(map[string]int)
	e["aaa"] = 1
	e["www"] = 2
	fmt.Println(e)

	//map第二个元素是切片的
	var d map[string][]int
	d = make(map[string][]int)
	aa := make([]int, 6)
	aa[0] = 111
	aa[1] = 111
	d["a"] = aa
	fmt.Println(d)

	//map第二个元素是切片的
	var r map[string][]int
	r = make(map[string][]int)
	tt := make([]int, 6)
	tt = append(tt, 1222)
	tt = append(tt, 1111)
	r["a"] = tt
	fmt.Println(r)

	//map第二个元素是切片的
	c := map[string][]int{"a": {1, 2, 3}, "b": {3, 3, 3, 3}}
	fmt.Println(c)
}
