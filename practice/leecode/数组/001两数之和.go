package main

import "fmt"

func main() {

	target := 6
	shuzu := []int{1, 2, 5, 6, 7, 9}

	res := two(target, shuzu)

	fmt.Println(res)
}

func two(target int, shuzu []int) (res []int) {
	for i, v := range shuzu {
		for j := i + 1; j < len(shuzu)-1; j++ {
			if v+shuzu[j] == target {
				//res = append(res, i, j)
				//fmt.Println(res)

				//fmt.Println([]int{i, j})

				return []int{i, j}
			}
		}
	}

	return nil
}
