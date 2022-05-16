package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("子协程可以打印了,因为定时器的时间到了")
	}()

	//关闭定时器
	timer.Stop()
}
