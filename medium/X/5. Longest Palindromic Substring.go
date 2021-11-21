package X

import (
	"math"
)

func longestPalindrome(s string) string {
	mem := make([][]bool, len(s))
	maxVal := math.MinInt32
	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		mem[i] = make([]bool, len(s))
		mem[i][i] = true
	}

	for j := 1; j < len(s); j++ {
		for i := 0; i < len(s)-j; i++ {
			k := i + j
			if s[i] == s[k] && (k == i+1 || mem[i+1][k-1]) {
				mem[i][k] = true
				if k-i+1 > maxVal {
					maxVal = k - i + 1
					start, end = i, k
				}
			}
		}
	}

	return s[start : end+1]
}
