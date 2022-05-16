package main

import "fmt"

type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	//此处会进行调用string方法  死循环了
	return fmt.Sprintf("print: %v", c)
}

func main() {
	c := &ConfigOne{}
	c.String()
}

//runtime: goroutine stack exceeds 1000000000-byte limit
//runtime: sp=0xc0201615b8 stack=[0xc020160000, 0xc040160000]
//fatal error: stack overflow

//参考答案及解析：
//运行时错误。如果类型实现 String() 方法，当格式化输出时会自动使用 String() 方法。上面这段代码是在该类型的 String() 方法内使用格式化输出，导致递归调用，最后抛错。
