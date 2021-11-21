package easy

import (
	"fmt"
	"strconv"
)

func validWordAbbreviation(word string, abbr string) bool {
	i := 0
	vi := 0
	for ; i < len(abbr); {
		a := abbr[i]
		if a >= 'a' && a <= 'z' || a >= 'A' && a <= 'Z' {
			if vi >= len(word) || abbr[i] != word[vi] {
				return false
			}

			vi++
			i++

			continue
		}

		p := i
		for ; i < len(abbr) && abbr[i] >= '0' && abbr[i] <= '9'; i, vi = i+1, vi+1 {
		}

		inc, _ := strconv.Atoi(abbr[p:i])

		if inc == 0 || fmt.Sprintf("%d", inc) != abbr[p:i] {
			return false
		}

		if inc > 0 {
			vi += inc - (i - p)
		}
	}

	return vi == len(word)
}
