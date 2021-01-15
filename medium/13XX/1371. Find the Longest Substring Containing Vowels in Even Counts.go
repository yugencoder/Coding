package t13XX

import (
	"strings"
)

func FindTheLongestSubstring(s string) int {
	res := 0
	currentSum := 0

	valueMap := map[int]int{0: -1}

	for i := range s {
		currentSum ^= 1 << (strings.IndexByte("aeiou", s[i]) + 1) >> 1

		if _, ok := valueMap[currentSum]; !ok {
			valueMap[currentSum] = i
			continue
		}

		res = max(res, i-valueMap[currentSum])
	}

	return res
}
