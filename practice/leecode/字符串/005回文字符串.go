package main

import "fmt"

func main() {
	fmt.Println(badLongestPalindrome("babad"))
}

func badLongestPalindrome(s string) string {
	runes := []rune(s)
	fmt.Println(runes)
	n := len(runes)
	if n <= 1 {
		return s
	}

	maxLen := 0
	maxL, maxR := 0, 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isPalindrome(runes[i:j+1]) && j-i+1 > maxLen {
				maxLen = j - i + 1
				maxL, maxR = i, j
			}
		}
	}
	return string(runes[maxL : maxR+1]) // 注意对空字符串，转换为 "\u0000"
}

// O(N) 判断是否为回文字符串
func isPalindrome(runes []rune) bool {
	l, r := 0, len(runes)-1
	for l <= r {
		if runes[l] != runes[r] {
			return false
		}
		l++
		r--
	}
	return true
}
