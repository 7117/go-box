package main

import (
	"github.com/xtaci/kcp-go"
	"time"
)

func main() {
	kcpconn, err := kcp.DialWithOptions("localhost:10000", nil, 10, 3)
	if err != nil {
		panic(err)
	}

	_, _ = kcpconn.Write([]byte("hello kcp"))
	time.Sleep(time.Second)
}
