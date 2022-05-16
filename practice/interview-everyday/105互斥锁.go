package main

import (
	"fmt"
	"sync"
)

type MyMutex struct {
	count int
	sync.Mutex
}

func main() {
	var mu MyMutex
	mu.Lock()
	fmt.Println(mu.count)
	mu.count++
	fmt.Println(mu.count)
	mu.Unlock()
	fmt.Println(mu.count)
}

//0
//1
//1
