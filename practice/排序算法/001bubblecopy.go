package main

import "fmt"

func main() {
	values := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	fmt.Println(values)
	BubbleAsortCopy(values)
}

func BubbleAsortCopy(values []int) {
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] < values[j] {
				values[j], values[i] = values[i], values[j]
			}
		}
	}
	fmt.Println(values)
}
