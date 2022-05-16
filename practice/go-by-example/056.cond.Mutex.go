package main

import (
	"fmt"
	"sync"
)

var lockera = new(sync.Mutex)

var conda = sync.NewCond(lockera)

// WaitGroup控制节奏的
var waitgroupa sync.WaitGroup

func test(x int) {
	conda.L.Lock()
	fmt.Println("test", x)
	defer func() {
		conda.L.Unlock()  //释放锁
		waitgroupa.Done() //减一个协程
	}()
}

func main() {

	for i := 0; i < 5; i++ {
		go test(i)
		waitgroupa.Add(1)
	}

	waitgroupa.Wait() //保证所有协程执行完毕
}

// test 0
// test 4
// test 3
// test 1
// test 2
