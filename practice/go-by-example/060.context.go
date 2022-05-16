package main

import (
	"context"
	"log"
	"time"
)

// 使用 Context 控制多个 goroutine
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "task 1")
	go watch(ctx, "task 2")
	go watch(ctx, "task 3")

	time.Sleep(3 * time.Second)
	log.Println("可以通知任务结束")
	cancel()
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			log.Println(name + " 任务即将要退出了...")
			return
		default:
			log.Println(name + " goroutine 继续处理任务中...")
			time.Sleep(2 * time.Second)
		}
	}
}
