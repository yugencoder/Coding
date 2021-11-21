package _7XX

import (
	"math"
	"sort"
)

func minOperations(nums1 []int, nums2 []int) int {
	sum1 := 0
	sum2 := 0

	for _, n := range nums1 {
		sum1 += n
	}

	for _, n := range nums2 {
		sum2 += n
	}

	n1 := nums1
	n2 := nums2
	if sum1 > sum2 {
		n1 = nums2
		n2 = nums1
	}

	sort.Ints(n1)
	sort.Slice(n2, func(i, j int) bool {
		return n2[i] > n2[j]
	})

	p1, p2 := 0, 0
	diff := int(math.Abs(float64(sum1 - sum2)))
	count := 0

	for ; p1 < len(n1) || p2 < len(n2); count++ {
		if diff <= 0 {
			break
		}

		if p1 == len(n1) {
			diff -= n2[p2] - 1
			p2++
			continue

		}

		if p2 == len(n2) || (6-n1[p1]) > (n2[p2]-1) {
			diff -= 6 - n1[p1]
			p1++
		} else {
			diff -= n2[p2] - 1
			p2++
		}

	}

	if diff <= 0 {
		return count
	}

	return -1
}
