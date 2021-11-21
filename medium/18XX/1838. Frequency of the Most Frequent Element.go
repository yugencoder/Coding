package _8XX

import (
	"math"
	"sort"
)

func maxFrequency(nums []int, k int) int {
	res := math.MinInt32
	currSum := 0

	i := 0

	sort.Ints(nums)

	for j, n := range nums {
		currSum += n

		for ; k+currSum < n*(j-i+1) && i < j; i++ {
			currSum -= nums[i]
		}
		res = max(res, j-i+1)

	}

	return max(0, res)
}
