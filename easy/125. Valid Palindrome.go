package easy

import (
	"strings"
)

func IsPalindrome(s string) bool {
	strippedR := []rune{}

	for _, c := range s {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c >= '0' && c <= '9' {
			newC := strings.ToLower(string(c))
			strippedR = append(strippedR, rune(newC[0]))
		}
	}

	s1 := string(strippedR)
	return strings.EqualFold(s1, string(rev(strippedR)))
}

func rev(runes []rune) []rune {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return runes
}
