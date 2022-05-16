package main

import "fmt"

type Tattoo struct {
	name string
}

func (p *Tattoo) print() {
	fmt.Println("name:", p.name)
}

type printer interface {
	print()
}

func main() {
	d1 := Tattoo{"one"}
	d1.print()

	var in printer = &Tattoo{"two"}
	in.print()
}

//name: one
//name: two
