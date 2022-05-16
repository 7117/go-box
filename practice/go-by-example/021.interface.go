package main

import "fmt"

type Animal interface {
	Fly()
	Run()
}

// 此处不需要写  interface
// 因为Birda实现了Animal的所有方法
// 所以就默认与animal进行了连接了！
type Birda struct {
}

func (Birda Birda) Fly() {
	fmt.Print("Birda is flying")
}

func (Birda Birda) Run() {
	fmt.Print("Birda is runing")
}

func main() {

	Birda := new(Birda)

	Birda.Fly()
	Birda.Run()

}
