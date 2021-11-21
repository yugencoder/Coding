package _8XX

import "strings"

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	words1 := strings.Split(sentence1, " ")
	words2 := strings.Split(sentence2, " ")

	idx1, idx2 := 0, 0
	extraSentence := false

	for ; idx1 < len(words1) && idx2 < len(words2); {
		if words1[idx1] == words2[idx2] {
			idx1++
			idx2++

		} else if extraSentence {
			return false
		} else {
			extraSentence = true

			left := false
			if len(words1) > len(words2) {
				left = true
			}

			for ; idx1 < len(words1) && idx2 < len(words2) && (len(words1)-idx1) != (len(words2)-idx2); {
				if left {
					idx1++
				} else {
					idx2++
				}
			}
		}
	}

	if ((idx1 == len(words1) && idx2 != len(words2)) || (idx1 != len(words1) && idx2 == len(words2))) && !extraSentence {
		return true
	}

	return idx1 == len(words1) && idx2 == len(words2)
}
