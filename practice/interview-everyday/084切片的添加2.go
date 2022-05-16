package main

import "fmt"

func main() {
	a := []int{0, 1, 2}
	s := a[1:2]
	fmt.Println("s1:", s)

	s[0] = 11
	s = append(s, 12)
	s = append(s, 13)
	s[0] = 21

	fmt.Println(a)
	fmt.Println(s)
}

//s1: [1]
//[0 11 12]
//[21 12 13]
//https://mp.weixin.qq.com/s/3qguB_V6mwPl-G2q-TjnfA
//定义是nil切片  nil切片是nil
//初始化赋值是空切片  空切片是【】
