package main

import "fmt"

type TUY struct {
	n int
}

func main() {
	ts := [2]TUY{}
	for i := range ts[:] {
		switch t := &ts[i]; i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ")
		}
	}
	fmt.Print(ts)
}

//9 [{3} {9}]
