package X

import "math"

func maxArea(height []int) int {
	res := math.MinInt32

	for i, h := range height {
		for j := len(height) - 1; j >= 0 && i < j; j-- {
			if height[j] >= h {
				res = max(res, (j-i)*h)
				break
			}

			if res > (j-i+1)*h {
				break
			}

			res = max(res, (j-i)*min(h, height[j]))
		}
	}

	return res
}
