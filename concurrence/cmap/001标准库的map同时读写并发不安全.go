package main

import "fmt"

func main() {
	testMapReadWriteDiffKey()
}

func testMapReadWriteDiffKey() {
	//fatal error: concurrent map read and map write
	m := make(map[int]int)
	go func() {
		for {
			m[100] = 100
		}
	}()
	go func() {
		for {
			read := m[12]
			fmt.Println(read)
		}
	}()
	select {}
}
