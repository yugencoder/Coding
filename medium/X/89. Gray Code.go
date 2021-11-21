package X

import "math"

func grayCode(n int) []int {
	len := int(math.Pow(2, float64(n)))
	res := make([]int, len)
	res[0] = 0
	res[1] = 1
	if n == 1 {
		return res
	}
	xor := 1
	count := 1
	for i := 2; i < len; i++ {
		if i&(i-1) == 0 {
			xor <<= 1
			res[i] = res[i-1] ^ xor
			count = 1
		} else {
			res[i] = res[i-1] ^ count&(^(count-1))
			count++
		}

	}

	return res
}
