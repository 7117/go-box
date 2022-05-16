package main

import (
	"fmt"
	"strings"
)

//编写一个函数来查找字符串数组中的最长公共前缀。
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	firstStr := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], firstStr) != 0 {
			firstStr = firstStr[:len(firstStr)-1]
		}
	}
	if len(firstStr) > 0 {
		return firstStr
	}
	return ""
}

func main() {
	s := []string{"flower", "flow", "floght"}
	fmt.Println(longestCommonPrefix(s))
}
