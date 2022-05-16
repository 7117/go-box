package main

type TTY []int

func FFT(t TTY) {}

func main() {
	//Unnamed Type 是基于已有的 Named Type 组合一起的类型
	//因为是【】int不是named type
	var q []int
	FFT(q)
}
