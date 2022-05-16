package main

import (
	"fmt"
	// socket包  操作tcp的
	"net"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func ProcessInfo(conn net.Conn) {
	// 缓存区 字节数组
	buf := make([]byte, 1024)
	defer conn.Close()

	for {
		numOfBytes, err := conn.Read(buf)
		if err != nil {
			break
		}
		if numOfBytes != 0 {
			fmt.Printf("received mess %s:", string(buf))
		}
		// 字符串数组  使用string把字节数组转化为字符串数组
		fmt.Printf("received:%s\n", string(buf))
	}
}

func main() {
	listen_socket, err := net.Listen("tcp", "127.0.0.1:8080")
	checkError(err)
	defer listen_socket.Close()

	fmt.Println("waiting")
	for {
		// 有人连接
		conn, err := listen_socket.Accept()
		checkError(err)
		go ProcessInfo(conn)
	}

}
