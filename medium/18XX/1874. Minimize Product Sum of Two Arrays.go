package _8XX

import (
	"sort"
)

func minProductSum(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] > nums2[j]
	})

	res := 0

	//fmt.Println(nums1)
	//fmt.Println(nums2)
	for i := 0; i < len(nums1); i++ {
		res += nums1[i] * nums2[i]
	}

	return res
}
