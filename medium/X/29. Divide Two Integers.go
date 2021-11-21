package X

import "math"

func divide(dividend int, divisor int) int {
	powerOfTwos := []int{}
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	sign := 1
	if dividend < 0 {
		sign = -1
		dividend *= -1
	}

	if divisor < 0 {
		sign *= -1
		divisor *= -1
	}

	res := 0

	power := 1
	for i := divisor; i <= dividend; i = i + i {
		powerOfTwos = append(powerOfTwos, i)
		if i > divisor {
			power <<= 1
		}
	}

	for i := len(powerOfTwos) - 1; dividend >= divisor; i-- {
		if powerOfTwos[i] <= dividend {
			dividend -= powerOfTwos[i]
			res += power
		}
		power >>= 1
	}

	return res * sign
}
