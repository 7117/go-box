package main

import "fmt"

func main() {

	eeeee := tolower("ADSED")
	fmt.Println("结果：", eeeee)

}

func tolower(aaa string) string {

	charDiff := 'a' - 'A'

	newstring := ""

	for _, r := range []int32(aaa) {
		if r < 'A' && r > 'Z' {
			newstring += string(r)
			continue
		}

		newstring += string(r + charDiff)
		fmt.Println(newstring)

	}

	//panic("")

	fmt.Println(newstring)

	return newstring
}
