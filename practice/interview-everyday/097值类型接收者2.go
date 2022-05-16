package main

import "fmt"

type Foosss struct {
	val int
}

func (f *Foosss) Inc(inc int) {
	f.val += inc
}

func main() {
	f := &Foosss{}
	f.Inc(100)
	fmt.Println(f.val) // 100
}

//100
