package main

func main() {
	defer func() {
		recover()
	}()
	panic(1)
}

//recover() 必须在 defer() 函数中直接调用才有效。
