package _7XX

func minimumLength(s string) int {
	res := len(s)
	i, j := 0, len(s)-1

	for i < j {
		iIdx := 1
		jIdx := 1
		if s[i] == s[j] {
			i++
			j--
			for i < len(s) && s[i] == s[i-1] {
				i++
				iIdx++
			}
			for j >= 0 && s[j] == s[j+1] {
				j--
				jIdx++
			}
			res = res - (iIdx + jIdx)

		} else {
			return res
		}
	}

	return max(res,0)
}
