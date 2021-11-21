package easy

import "math"

func Reverse(x int) int {
	res := 0
	isNeg := false
	if x < 0 {
		x = x * -1
		isNeg = true

	}

	for ; x > 0; {
		if (res > math.MaxInt32/10) || (res == math.MaxInt32/10 && x%10 > (math.MaxInt32-1)%10) {
			return 0
		}
		if isNeg {
			if (res > math.MaxInt32/10) || (res == math.MaxInt32/10 && x%10 > (math.MaxInt32-1)%10) {
				return 0
			}
		}

		res = res*10 + x%10
		x = x / 10
	}

	if isNeg {
		return -res
	}

	return res
}

// 321
// 1*10+2
