package main

import (
	"fmt"
	"sync"
	"time"
)

type datatwo struct {
	*sync.Mutex
}

func (d datatwo) test(s string) {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 2; i++ {
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	d := datatwo{new(sync.Mutex)}

	go func() {
		defer wg.Done()
		d.test("read")
	}()

	go func() {
		defer wg.Done()
		d.test("write")
	}()

	wg.Wait()
}

//write 0
//write 1
//read 0
//read 1
