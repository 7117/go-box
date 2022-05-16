package main

import (
	"fmt"
	"time"
)

func main() {

	var m = [...]int{1, 2, 3}

	for i, v := range m {
		go func() {
			fmt.Println(i, v)
		}()
	}

	for i, v := range m {
		go func(i, v int) {
			fmt.Println(i, v)
		}(i, v)
	}

	time.Sleep(time.Second * 1)
}

//PS D:\goitem\go-collection\practice\interview-everyday> go run .\029.go
//2 3
//2 3
//2 3
//2 3
//0 1
//1 2
