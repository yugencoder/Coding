package _7XX

func largestMerge(word1 string, word2 string) string {
	var res string

	w1, w2 := 0, 0

	for ; w1 < len(word1) && w2 < len(word2); {
		lesser := getGreater(word1[w1:], word2[w2:])
		res = res + string(lesser[0])
		if lesser == word1[w1:] {
			w1++
		} else {
			w2++
		}
	}

	if w2 < len(word2) {
		return res + word2[w2:]
	} else if w1 < len(word1) {
		return res + word1[w1:]
	}

	return res
}

func getGreater(s1 string, s2 string) string {
	l := min(len(s1), len(s2))

	for i := 0; i < l; i++ {
		if s1[i] > s2[i] {
			return s1
		} else if s1[i] < s2[i] {
			return s2
		}
	}

	if l == len(s1) {
		return s2
	}

	return s1
}
