package main

import "fmt"

type ConfigOneE struct {
	Daemon string
}

func (c *ConfigOneE) String() string {
	return fmt.Sprintf("print: %v", c)
}

func main() {
	c := &ConfigOneE{}
	c.String()
}

//无限递归循环，栈溢出。
//知识点：类型的 String() 方法。如果类型定义了 String() 方法，使用 Printf()、Print() 、 Println() 、 Sprintf() 等格式化输出时会自动使用 String() 方法。
