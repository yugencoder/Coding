package _7XX

import "math"

func minElements(nums []int, limit int, goal int) int {
	sum := 0

	for _, n := range nums {
		sum += n
	}

	diff := int(math.Abs(float64(goal - sum)))

	if diff % limit ==0 {
		return diff/limit
	}
	return diff/limit+1
}
