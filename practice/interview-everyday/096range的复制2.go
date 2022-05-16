package main

import "fmt"

type Pointtt struct {
	x, y int
}

func main() {
	s := []*Pointtt{
		{1, 2},
		{3, 4},
	}
	for _, p := range s {
		p.x, p.y = p.y, p.x
	}
	fmt.Println(*s[0])
	fmt.Println(*s[1])
}

//{2 1}
//{4 3}
