package _8XX

import (
	"math"
	"sort"
)

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	diff := make([][]int, len(nums1))
	nums := make([][]int, len(nums1))
	sum := 0

	for i, _ := range nums1 {
		nums[i] = make([]int, 2)
		diff[i] = make([]int, 2)

		diff[i][0] = int(math.Abs(float64(nums1[i] - nums2[i])))
		sum = (sum + diff[i][0]) % 1000000007
		diff[i][1] = i

		nums[i][0] = nums1[i]
		nums[i][1] = i
	}

	sort.Slice(diff, func(i, j int) bool {
		return diff[i][0] > diff[j][0]
	})

	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})

	minDiff := math.MaxInt32
	oldDiff := 0
	for i, d := range diff {
		gIdx := getGreaterThanEqualTo(nums2[d[1]], nums, 0, len(nums1)-1)
		lIdx := getLesserThanEqualTo(nums2[d[1]], nums, 0, len(nums1)-1)

		newDiff := min(int(math.Abs(float64(nums2[d[1]]-nums1[gIdx]))), int(math.Abs(float64(nums2[d[1]]-nums1[lIdx]))))

		if i < len(diff)-1 && int(math.Abs(float64(newDiff-d[0]))) >= diff[i+1][0] {
			return (sum + newDiff - d[0]) % 1000000007
		} else {
			minDiff = min(minDiff, newDiff)
			if minDiff == newDiff {
				oldDiff = d[0]
			}
		}

	}

	return (sum - oldDiff + minDiff) % 1000000007
}
