package main

import (
	"fmt"
	"time"
)

func testdddd() {
	fmt.Print("testdddd")
}

func main() {
	go testdddd()

	time.Sleep(2)
}
