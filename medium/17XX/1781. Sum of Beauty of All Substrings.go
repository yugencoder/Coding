package _7XX

import "math"

func beautySum(s string) int {
	res := 0
	for i := range s {
		freq := make([]int, 26)
		minV := math.MaxInt32
		maxV := math.MinInt32

		for j := i; j < len(s); j++ {
			freq[s[j]-'a']++

			maxV = max(maxV, freq[s[j]-'a'])
			minV = math.MaxInt32

			for _, v := range freq {
				if v == 0 {
					continue
				}
				minV = min(minV, v)
			}

			res += maxV - minV

		}
	}

	return res
}
