package main

func main() {

	//创建通道
	//出现死锁是因为没有goroutine在写
	//因此，select语句被永远阻塞
	c := make(chan int)
	select {
	case <-c:
	}
}
