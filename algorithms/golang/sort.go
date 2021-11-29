package golang

import (
	"fmt"
	"sort"
)

func sortAscendingInt(nums []int) {
	sort.Ints(nums)

	fmt.Println(nums)
}

func sortDescendingInt(nums []int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	fmt.Println(nums)
}
