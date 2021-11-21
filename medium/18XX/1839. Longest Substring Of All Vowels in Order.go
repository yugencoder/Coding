package _8XX

import "math"

//Input: word = "aeiaaioaaaaeiiiiouuuooaauuaeiu"
//Output: 13
//Explanation: The longest beautiful substring in word is "aaaaeiiiiouuu" of length 13.

func longestBeautifulSubstring(word string) int {
	maxLen := math.MinInt32
	currLen := 0

	nextVowel := map[string]string{
		"a": "e",
		"e": "i",
		"i": "o",
		"o": "u",
	}
	//vowelMap := map[string]bool{
	//	"a": true, "e": true, "i": true, "o": true, "u": true,
	//}

	possibleMatch := false
	prev := "-1"
	curr := "-1"

	for i := range word {
		curr = string(word[i])

		if curr == "a" && prev != "a" {
			possibleMatch = true
			currLen = 0
		}

		if !possibleMatch {
			prev = curr

			continue
		}

		if (curr == prev || curr == "a") || (curr == nextVowel[prev]) {
			currLen++
			if curr == "u" {
				maxLen = max(maxLen, currLen)
			}
		} else {
			possibleMatch = false
		}

		prev = curr
	}

	return max(0, maxLen)
}

