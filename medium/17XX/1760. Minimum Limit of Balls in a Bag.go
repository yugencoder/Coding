package _7XX

import "math"

func minimumSize(nums []int, maxOperations int) int {
	start := 0
	end := math.MaxInt32

	for;start < end ;{
		mid := start + (end - start)/2

		ops := 0
		for _, n := range nums{
			ops += (n-1)/mid
		}

		if ops > maxOperations{
			start = mid+1
		}else{
			end = mid
		}
	}

	return start
}
