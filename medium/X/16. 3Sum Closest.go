package X

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	res := 0
	minDiff := math.MaxInt32
	sort.Ints(nums)

	for i := range nums {
		start := i + 1
		end := len(nums)-1
		for start < end {
			diff := findDiff(nums[i]+nums[start]+nums[end], target)
			if diff < minDiff {
				minDiff = diff
				res = nums[i] + nums[start] + nums[end]
			}

			if findDiff(nums[i]+nums[start+1]+nums[end], target) < diff {
				start++
			} else {
				end--
			}
		}
	}

	return res
}

func findDiff(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
