package _7XX

import "fmt"

func canChoose(groups [][]int, nums []int) bool {
	k := 0

	for i := 0; i < len(nums); {
		if nums[i] == groups[k][0] {
			if isMatched(groups[k], nums[i:]) {
				i = i + len(groups[k])
				k++

				if k == len(groups) {
					return true
				}
				continue
			}
		}
		i++
	}

	return false
}

func isMatched(first, second []int) bool {
	if len(first) > len(second) {
		return false
	}

	firstStr := fmt.Sprintf("%v", first)
	secondStr := fmt.Sprintf("%v", second[:len(first)])

	return firstStr == secondStr
}
