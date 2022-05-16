package main

import "fmt"

func main() {
	fmt.Println(jisuan([]int{1, 3, 5}, []int{2, 4, 6}))
}

func jisuan(num1 []int, num2 []int) float64 {
	res := merge(num1, num2)

	n := len(res)

	if n == 0 {
		return -1
	}
	if n%2 == 0 {
		fmt.Println(res[n/2-1])
		fmt.Println(res[n/2])
		return float64(res[n/2-1]+res[n/2]) / 2
	}
	return float64(res[n/2])
}

func merge(num1 []int, num2 []int) []int {
	n1, n2 := len(num1), len(num2)
	res := make([]int, 0, n1+n2)
	i, j := 0, 0

	for i < n1 && j < n2 {
		switch {
		case num1[i] < num2[j]:
			res = append(res, num1[i])
			i++
		case num1[i] > num2[j]:
			res = append(res, num2[j])
			j++
		default:
			res = append(res, num1[i], num2[j])
			i++
			j++
		}
	}

	if i < n1 {
		res = append(res, num1[i:]...)
	}

	if j < n2 {
		res = append(res, num2[j:]...)
	}

	fmt.Println(res)

	return res
}
