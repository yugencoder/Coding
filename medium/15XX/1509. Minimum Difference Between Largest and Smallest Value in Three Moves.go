package _5XX

import (
	"math"
	"sort"
)

func minDifference(nums []int) int {
	res := math.MaxInt32
	wLen := len(nums) - 4
	if wLen <= 0 {
		return 0
	}

	sort.Ints(nums)

	for j := len(nums) - 1; j >= 0; j -= 1 {
		if j-wLen >= 0 {
			res = min(res, nums[j]-nums[j-wLen])
		}
	}

	return res
}
