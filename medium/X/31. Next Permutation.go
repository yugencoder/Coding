package X

import (
	"sort"
)

func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}

	i := len(nums) - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}

	if i < 0 {
		sort.Ints(nums)
		return
	}


	for j := len(nums)-1; j > i; j-- {
		if nums[j] > nums[i]{
			nums[i], nums[j] = nums[j], nums[i]
			break
		}
	}

	sort.Ints(nums[i+1:])
}
