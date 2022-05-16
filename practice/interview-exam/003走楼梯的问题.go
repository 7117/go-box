package main

import "fmt"

func main() {
	fmt.Println(climbStairs1(3)) // 3
}

//通过枚举进行总结出规律来即可
func climbStairs1(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs1(n-1) + climbStairs1(n-2)
}
