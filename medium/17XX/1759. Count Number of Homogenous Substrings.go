package _7XX

func countHomogenous(s string) int {
	res := 0
	next := int32(-1)
	l := 1

	for i, c := range s {
		next = int32(-1)

		if i+1 < len(s) {
			next = int32(s[i+1])
		}

		if next != c {
			res += l * (l + 1) / 2 % 1_000_000_007
			l = 1
		} else {
			l++
		}
	}

	return res
}
