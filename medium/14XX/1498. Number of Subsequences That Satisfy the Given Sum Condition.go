package _14XX

import (
	"sort"
)

func numSubseq(nums []int, target int) int {
	var res int
	const MOD = 1e9 + 7

	exps := make([]int, len(nums))
	exps[0] = 1

	for i := 1; i < len(exps); i++ {
		exps[i] = (2 * exps[i-1]) % MOD
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for i, n := range nums {
		if n > target {
			continue
		}

		for j, c := i+1, 0; j < len(nums); j, c = j+1, c+1 {
			if n+nums[j] <= target {
				res += exps[len(nums)-(i+1)] - (exps[c] - 1) - 1
				break
			}
		}

		if 2*n <= target {
			res++
		}
		res = res % MOD
	}

	return res
}
