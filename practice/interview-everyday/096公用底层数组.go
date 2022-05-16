package main

import "fmt"

func get() []byte {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0])
	return raw[:3]
}

func main() {
	data := get()
	fmt.Println(len(data), cap(data), &data[0])
}

//10000 10000 0xc000112000
//3 10000 0xc000112000
