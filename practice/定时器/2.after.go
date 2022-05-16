package main

import (
	"fmt"
	"time"
)

func main() {
	//定时2秒,2秒后产生一个事件,往channel里面写内容
	<-time.After(3 * time.Second)
	fmt.Println("两秒时间到")
}