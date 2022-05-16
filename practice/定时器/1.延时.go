package main

import (
	"fmt"
	"time"
)

func main() {
	//time.NewTimer时间到了,只响应一次
	//创建一个定时器,设置时间为2s,2s后,往time通道写内容(当前时间)
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("当前时间: ", time.Now())

	//2s后,往timer.c写数据,有数据后,就可以读取
	t := <-timer.C //channel没有数据前后阻塞
	fmt.Println("t = ", t)
}
