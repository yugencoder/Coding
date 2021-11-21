package _7XX

import "sort"

func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)

	res := 1
	for _, c := range coins {
		if c > res {
			return res
		}
		res += c
	}
	return res
}
