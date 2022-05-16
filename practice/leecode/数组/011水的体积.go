package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxArea1([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea1(h []int) int {
	n := len(h)

	if n < 1 {
		return 0
	}

	max_area := 0
	for i := 1; i < n; i++ {
		for j := i + 1; j < n; j++ {
			w := j - i

			h := math.Min(float64(h[i]), float64(h[j]))

			max_area = int(math.Max(float64(max_area), float64(float64(w)*h)))
		}
	}

	return max_area
}
