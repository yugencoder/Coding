package X

import "math"

func myAtoi(s string) int {
	sign := 1
	needsDigit := false
	num := 0

	for _, c := range s {
		if needsDigit && !(c >= '0' && c <= '9') {
			break
		}

		if c == ' ' {
			continue
		}

		if c == '+' || c == '-' {
			if c == '-' {
				sign = -1
			}

			needsDigit = true
			continue
		}

		if c >= '0' && c <= '9' {
			needsDigit = true
			digit := int(c - '0')

			if num == math.MaxInt32/10 && digit > math.MaxInt32%10 || num > math.MaxInt32/10 {
				if sign < 0 {
					return math.MinInt32
				}
				return math.MaxInt32
			}

			num = num*10 + digit
		} else {
			break
		}
	}

	return sign * num
}
