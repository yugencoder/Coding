package aiq

import (
	"regexp"
	"strings"
	"unicode"
)

func howManyWords(s string) int {
	words := strings.Fields(s)
	res := 0

	for _, w := range words {
		arr := strings.Split(w, "-")
		if len(arr) == 1 {
			if isWord(arr[0]) {
				res++
			}
		} else if len(arr) == 2 {
			if isWord(arr[0]) && isWord(arr[1]) {
				res++
			}
		}
	}

	return res

}

func howManyWords2(s string) int {
	words := strings.Fields(s)
	res := 0
	expr := `^[A-Za-z]+(?:-[A-Za-z]+)*[.,;:?!]?$`
	wordMatcher := regexp.MustCompile(expr)

	for _, w := range words {
		if wordMatcher.MatchString(w){
			res++
		}
	}

	return res

}

func isWord(s string) bool {
	for i, c := range s {
		if !(unicode.IsLetter(c) || (i > 0 && i == len(s)-1 && (c == '!' || c == '?' || c == '.' || c == ','))) {
			return false
		}
	}

	return true
}
