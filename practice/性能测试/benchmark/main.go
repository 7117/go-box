package benchmark

//被测试函数
func sum(a, b int) int {
	return a + b
}

//被测试函数
func sliceNoCap() {
	s := make([]int, 0) // 未设置 cap 值
	for i := 0; i < 10000; i++ {
		s = append(s, i)
	}
}

//被测试函数
func sliceWithCap() {
	s := make([]int, 0, 10000) // 预先设置 cap 值
	for i := 0; i < 10000; i++ {
		s = append(s, i)
	}
}

