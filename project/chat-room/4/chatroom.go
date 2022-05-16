package main

import (
	"fmt"
	// socket包  操作tcp的
	"log"
	"net"
	"os"
	"strings"
)

var onlineConns = make(map[string]net.Conn)
var messageQueue = make(chan string, 1000)
var quitChan = make(chan bool)
var logFile *os.File
var logger *log.Logger

const (
	LOG_DIRECTORY = "log.log" //记录错误日志的路径
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// 处理交互信息:把客户端的信息写入到队列channel中
func ProcessInfo(conn net.Conn) {
	// 缓存区 字节数组
	buf := make([]byte, 1024)
	defer func(conn net.Conn) {
		addr := fmt.Sprintf("%s", conn.RemoteAddr())
		delete(onlineConns, addr)

		for i := range onlineConns {
			fmt.Println(i)
		}

		fmt.Printf("aaaaa")

	}(conn)

	for {
		// 把conn里面的信息读取到buf里面  返回的是一个长度  一个错误
		numOfBytes, err := conn.Read(buf)

		if err != nil {
			panic(err)
		}
		if numOfBytes != 0 {
			message := string(buf[0:numOfBytes])
			// 写入到chan中
			messageQueue <- message
		} else {
			fmt.Println("为0")
		}
	}
}

// 对于chnnel中的数据进行处理
func ConsumeMessage() {
	for {
		select {
		case message := <-messageQueue:
			//对消息进行解析
			doProcessMessage(message)
		case <-quitChan:
			break
		}
	}
}

func doProcessMessage(aaa string) {
	// 进行分割
	contents := strings.Split(aaa, "#")

	if len(contents) > 1 {
		// 地址
		addr := contents[0]
		// 内容
		sendMessage := strings.Join(contents[1:], "#")
		// 对于IP的处理
		addr = strings.Trim(addr, " ")
		fmt.Println(sendMessage)

		//通过addr查看是否有链接对象
		//根据端口找到应该的客户端
		if conn, ok := onlineConns[addr]; ok {
			fmt.Println(sendMessage)
			_, err := conn.Write([]byte(sendMessage))
			if err != nil {
				fmt.Println("online conns send failure!")
				panic(err)
			}
		}
	} else {
		//走到这里 说明客户端调用list命令 查看系统当前链接ip
		contents := strings.Split(aaa, "*")
		if strings.ToUpper(contents[1]) == "LIST" {
			// 组装IP集合
			var ips string = ""
			for i := range onlineConns {
				ips = ips + "|" + i
			}

			if conn, ok := onlineConns[contents[0]]; ok {
				// 写入conn
				_, err := conn.Write([]byte(ips))
				if err != nil {
					fmt.Println("online conns send failure!")
				}
			}
		}
	}
}

func main() {
	// 指定log文件
	logFile, err := os.OpenFile(LOG_DIRECTORY, os.O_RDWR|os.O_CREATE, 0)

	if err != nil {
		fmt.Println("创建失败")
		os.Exit(-1)
	}

	defer logFile.Close()
	//利用go自带的log 将打开文件对象生成日志文件对象
	// 设置格式
	logger = log.New(logFile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)

	listen_socket, err := net.Listen("tcp", "127.0.0.1:8080")
	checkError(err)
	defer listen_socket.Close()

	fmt.Println("waiting")
	logger.Println("writing the log")

	// 消费消息队列
	// 现在channel没有数据就会堵塞
	go ConsumeMessage()

	for {
		fmt.Println("waiting2")

		// 有人连接
		conn, err := listen_socket.Accept()
		fmt.Printf("%v\n", conn)
		checkError(err)

		// 将conn连接对象存储到onlineonns的映射表中
		// 键值为IP
		// 值为conn对象
		// 客户端的IP
		addr := fmt.Sprintf("%s", conn.RemoteAddr())
		onlineConns[addr] = conn

		// 把所有连接的IP进行输出出来
		for i := range onlineConns {
			fmt.Println(i)
		}

		// 处理交互信息
		go ProcessInfo(conn)
	}

}
