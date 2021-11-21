package X

import "math"

// Better solution is to backtrack.
func subsets(nums []int) [][]int {
	res := [][]int{}

	for i := 0; i < int(math.Pow(2, float64(len(nums)))); i++ {
		comb := []int{}
		mask := i
		one := 1

		for j := 0; j < len(nums); j++ {
			if one&mask > 0 {
				comb = append(comb, nums[j])
			}
			one <<= 1
		}
		res = append(res, comb)
	}

	return res
}
