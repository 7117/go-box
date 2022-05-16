package main

func GetValue() interface{} {
	return 1
}

func main() {
	i := GetValue()
	//看本质是什么样的类型
	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}
