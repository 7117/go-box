package main

import (
	"fmt"
	// socket包  操作tcp的
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Printf("error is:%s", err.Error())
		os.Exit(1)
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	checkError(err)
	defer conn.Close()

	conn.Write([]byte("hello iam client"))

	fmt.Println("has sent mess")

}
