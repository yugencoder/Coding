package X

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -1 * nums[i]
		start := i + 1
		end := len(nums) - 1

		for start < end {
			currSum := nums[start] + nums[end]
			if currSum == target {
				res = append(res, []int{nums[i], nums[start], nums[end]})

				for start < end && nums[start] == nums[start+1]{start++}
				for start < end && nums[end] == nums[end-1]{end--}

				start++
				end--
			} else if currSum < target {
				start++
			} else {
				end--
			}
		}
	}

	return res
}
