package main

import "fmt"

type Fooss struct {
	val int
}

func (f Fooss) Inc(inc int) {
	f.val += inc
}

func main() {
	var f Fooss
	f.Inc(100)
	fmt.Println(f.val)
}

//0
