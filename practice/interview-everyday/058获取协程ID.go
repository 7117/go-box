package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
)

func GetGoroutineID() uint64 {
	b := make([]byte, 64)
	runtime.Stack(b, false)
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	fmt.Println(n)
	return n
}

func main() {
	GetGoroutineID()
}
