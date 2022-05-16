package main

import "fmt"

type Man struct {
	Name string
}

func main() {

	var Me map[string]*Man
	Me = make(map[string]*Man)
	Me["class"] = &Man{
		Name: "sun",
	}
	fmt.Println(Me)
	fmt.Println(fmt.Sprintf("%+v\n", *Me["class"]))
}

//map[class:0xc000050230]
//1
//main.Man{Name:"sun"}
