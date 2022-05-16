package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)
	//重置上面的那个无效
	flag := timer.Reset(1 * time.Second)
	fmt.Println(flag)
	<-timer.C
	fmt.Println("时间到")
}
