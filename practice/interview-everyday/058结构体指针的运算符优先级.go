package main

import "fmt"

type TTW struct {
	x int
	y *int
}

func main() {

	i := 20
	t := TTW{10, &i}

	p := &t.x

	*p++
	*p--

	t.y = p

	fmt.Println(*t.y)
}

//10
