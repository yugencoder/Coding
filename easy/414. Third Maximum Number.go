package easy

import (
	"math"
)

func thirdMax(nums []int) int {
	firstMax, secMax, thirdMax := math.MinInt32,math.MinInt32,math.MinInt32

	for _, n:= range nums{
		if n > firstMax{
			firstMax, secMax, thirdMax = n,firstMax, secMax
		}else if n > secMax && n != firstMax{
			secMax, thirdMax = n, secMax
		}else if n > thirdMax && n != firstMax && n != secMax{
			thirdMax = n
		}
	}

	if secMax <= thirdMax || len(nums) < 3{
		return firstMax
	}

	return thirdMax
}
