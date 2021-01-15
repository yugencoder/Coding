package _5XX

import "math"

func getLastMoment(n int, left []int, right []int) int {
	lMax := math.MinInt32
	rMin := math.MaxInt32

	for _, l := range left {
		lMax = max(lMax, l)
	}

	for _, r := range right {
		rMin = min(rMin, r)
	}

	if rMin == math.MaxInt32 {
		return lMax
	}

	return max(lMax, n-rMin)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
