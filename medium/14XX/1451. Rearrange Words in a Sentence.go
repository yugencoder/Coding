package _14XX

import (
	"sort"
	"strings"
)

func arrangeWords(text string) string {
	words := strings.Split(text, " ")
	words[0] = strings.ToLower(string(words[0][0])) + words[0][1:]

	sort.SliceStable(words, func(i, j int) bool {
		return len(words[i]) <= len(words[j])
	})

	words[0] = strings.ToUpper(string(words[0][0])) + words[0][1:]

	return strings.Join(words, " ")
}
