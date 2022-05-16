package main

//func (m map[string]string) Set(key string, value string) {
//	m[key] = value
//}
//
//func main() {
//	m := make(map[string]string)
//	m.Set("A", "One")
//}

type UserTT map[string]string

func (m UserTT) Set(key string, value string) {
	m[key] = value
}

func main() {
	m := make(UserTT)
	m.Set("A", "One")
}

//Unnamed Type 不能作为方法的接收者。昨天我们讲过 Named Type 与 Unamed Type 的区别
