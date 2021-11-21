package _14XX

import "math"

func longestSubarray2(nums []int) int {
	numsMap := map[int]int{}

	prev := 0
	oneCount := 0
	for i, n := range nums {
		if n == 1 {
			oneCount++
		} else {
			numsMap[prev] = oneCount
			oneCount = 0
			prev = i
		}
	}

	numsMap[prev] = oneCount

	maxV := math.MinInt32
	oneCount = 0
	for i, n := range nums {
		if n == 1 {
			oneCount++
		} else {
			maxV = max(maxV, oneCount+numsMap[i])
			oneCount = 0
		}

		if i == len(nums)-1 {
			maxV = max(maxV, oneCount)
			if maxV == len(nums) {
				maxV--
			}
		}
	}

	return max(maxV, 0)
}
