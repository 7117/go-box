package main

import (
	"fmt"
	"sync"
	"time"
)

type data struct {
	sync.Mutex
}

func (d data) test(s string) {
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
	var d data

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

//参考答案及解析：锁失效。将 Mutex 作为匿名字段时，相关的方法必须使用指针接收者，否则会导致锁机制失效。
//write 0
//read 0
//read 1
//write 1
