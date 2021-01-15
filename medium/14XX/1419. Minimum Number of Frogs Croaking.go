package _14XX

import "math"

func minNumberOfFrogs(croakOfFrogs string) int {
	croaks := map[rune]int{}
	nextMap := map[rune]rune{
		'c': 'r',
		'r': 'o',
		'o': 'a',
		'a': 'k',
	}

	startRune, endRune := 'c', 'k'
	frogs, res := 0, math.MinInt32

	for _, c := range croakOfFrogs {
		if c == startRune {
			frogs++

			res = max(res, frogs)
			croaks[nextMap[c]] += 1

			continue
		}

		if _, ok := croaks[c]; ok {
			if croaks[c] -= 1; croaks[c] == 0 {
				delete(croaks, c)
			}

			if c == endRune {
				frogs--
			} else {
				croaks[nextMap[c]] += 1
			}

		} else {
			return -1
		}
	}

	if len(croaks) != 0 {
		return -1
	}

	return res
}
