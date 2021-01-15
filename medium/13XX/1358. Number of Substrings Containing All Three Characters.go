package t13XX

func numberOfSubstrings(s string) int {
	res := 0
	charCounts := buildCountMap(s)

	for i := 0; i < len(s)-2;  {
		if end := hasSubString(i, s, charCounts); end != 0 {
			for ; i < end && ( i < 1 || (charCounts[end][rune(s[i-1])-'a']-charCounts[i][rune(s[i-1])-'a']) > 0); i++ {
				res += 1 + len(s) - end
			}
		} else {
			return res
		}
	}

	return res
}

func buildCountMap(s string) map[int][]int {
	charCounts := map[int][]int{}

	var as, bs, cs int
	for i, c := range s {
		charCounts[i] = []int{as, bs, cs}
		if c == 'a' {
			as++
		} else if c == 'b' {
			bs++
		} else if c == 'c' {
			cs++
		}
	}

	charCounts[len(s)] = []int{as, bs, cs}

	return charCounts
}

func hasSubString(start int, s string, charsCountMap map[int][]int) int {
	for i := start + 3; i <= len(s); i++ {
		if (charsCountMap[i][0]-charsCountMap[start][0]) >= 1 && (charsCountMap[i][1]-charsCountMap[start][1]) >= 1 && (charsCountMap[i][2]-charsCountMap[start][2]) >= 1 {
			return i
		}
	}

	return 0
}

