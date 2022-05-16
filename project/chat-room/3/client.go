package main

import (
	// 读输入输出流
	"bufio"
	"fmt"
	// socket包  操作tcp的
	"net"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// 写入数据的处理
func messagesend(conn net.Conn) {
	var input string

	for {
		// 这是在写入数据的时候的操作！
		// 读取终端是不是有数据
		reader := bufio.NewReader(os.Stdin)
		// 获取数据
		data, _, _ := reader.ReadLine()
		// 数据处理
		input = string(data)

		if strings.ToUpper(input) == "EXIT" {
			conn.Close()
			break
		}

		// 写入到连接
		_, err := conn.Write([]byte(input))

		if err != nil {
			conn.Close()
			fmt.Printf("fail:%s\n", err.Error())
			break
		}
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	checkError(err)
	defer conn.Close()

	// conn.Write([]byte("hello iam client"));
	go messagesend(conn)

	//主协程负责接收消息
	buf := make([]byte, 1024)
	for {
		fmt.Println("00000")

		numOfBytes, err := conn.Read(buf)
		checkError(err)

		/*结尾buf[0:numOfBytes]的原因是：numOfBytes是指接收的字节数 如果只用string(buf)
		可能会导致接收字符串中有其他之前接收的字符
		*/
		fmt.Println("receive server message content:" + string(buf[0:numOfBytes]))
	}

	fmt.Println("Client program end!")

}
