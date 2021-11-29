package IQ

import (
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

}

func isWord(s string) bool {
	for _, c := range s {
		if !unicode.IsLower(c) {
			return false
		}
	}

	return true
}
