package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"flower", "flow", "flight"}
	a := s[0]

	for k, _ := range s {
		for strings.Index(s[k], a) != 0 {
			a = a[:len(a)-1]
		}
	}

	fmt.Println(a)
}
