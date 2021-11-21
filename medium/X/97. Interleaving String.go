package X

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	i := 0
	j := 0
	matches := 0
	doubleMatches := 0
	for k := range s3 {
		hasMatch := false
		if i < len(s1) && j < len(s2) && s1[i] == s3[k] && s2[j] == s3[k] {
			doubleMatches++
			//hasMatch = true
			//someMap[s1] = map[string]bool{}
			//someMap[s1][s2] = isInterleave(s1[i+1:], s2[j:], s3[k+1:]) || isInterleave(s1[i:], s2[j+1:], s3[k+1:])
			//return someMap[s1][s2]
		}

		if i < len(s1) && s1[i] == s3[k] {
			i++
			hasMatch = true
			matches++
		}
		if j < len(s2) && s2[j] == s3[k] {
			hasMatch = true
			j++
			matches++
		}

		if !hasMatch {
			if i-doubleMatches < len(s1) && i-doubleMatches >= 0 && s1[i-doubleMatches] == s3[k] {
				i -= doubleMatches - 1
			} else if j-doubleMatches < len(s2) && j-doubleMatches > 0 && s2[j-doubleMatches] == s3[k] {
				j -= doubleMatches - 1
			} else {
				return false
			}
			doubleMatches = 0
		}

	}
	if i >= len(s1) && j >= len(s2) && matches == len(s3) {
		return true
	}

	return false
}
