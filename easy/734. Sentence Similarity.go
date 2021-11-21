package easy

func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
	if len(sentence1) != len(sentence2) {
		return false
	}

	similarPairsMap := map[string][]int{}
	for i, p := range similarPairs {
		similarPairsMap[p[0]] = append(similarPairsMap[p[0]], i)
		similarPairsMap[p[1]] = append(similarPairsMap[p[1]], i)
	}

	for i := 0; i < len(sentence1); i++ {
		if sentence1[i] == sentence2[i] {
			continue
		} else if hasCommon(similarPairsMap[sentence1[i]], similarPairsMap[sentence2[i]]) {
			return false
		}
	}

	return true
}

func hasCommon(v1, v2 []int) bool {
	for _, v := range v1 {
		for _, w := range v2 {
			if v == w {
				return true
			}
		}
	}
	return false
}
