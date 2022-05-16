package main

import "fmt"

const (
	Century = 100
	Decade  = 010
	Year    = 001
)

func main() {
	//100+2*8+2*1=118
	fmt.Println(Century + 2*Decade + 2*Year)
}

//118
