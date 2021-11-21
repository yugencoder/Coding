package X

func lengthOfLongestSubstring(s string) int {
	prevMap := map[rune]int{}
	start := 0
	res := 0

	for i, c := range s {
		if v, ok := prevMap[c]; ok {
			if v > start {
				start = v
			}
		}
		prevMap[c] = i
		res = max(res, i-start)
	}

	return res
}
