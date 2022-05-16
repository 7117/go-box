package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	temp := arr[0]
	left := make([]int, 0)
	right := make([]int, 0)

	for i := 0; i < len(arr); i++ {
		if arr[i] > temp {
			left = append(left, arr[i])
		} else if arr[i] < temp {
			right = append(right, arr[i])
		}
	}

	left = quickSort(left)
	right = quickSort(right)
	arr = append(left, temp)
	arr = append(arr, right...)

	return arr
}
func main() {
	arr := []int{3, 7, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2}
	res := quickSort(arr)

	fmt.Println(res)
}
