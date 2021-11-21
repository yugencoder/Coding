package X

import "strconv"

var mappings map[string]int

func numDecodings(s string) int {
	mappings = map[string]int{}
	if len(s) == 0 {
		return 1
	}

	return noOWays(0, s)
}

func noOWays(i int, s string) int {
	r := []rune(s)

	subStr := string(r[i:])
	if v := mappings[subStr]; v > 0 {
		return v
	}

	if i >= len(s) || len(r[i:]) == 0 {
		return 1
	}

	res := 0
	if string(r[i]) != "0" && len(r[i:]) >= 2 && getMapping(r[i:i+2]) {
		res += noOWays(i+2, s)
	}

	if len(r[i:]) >= 1 && getMapping(r[i:i+1]) {
		res += noOWays(i+1, s)
	}

	mappings[subStr] =  res

	return res
}

func getMapping(runes []rune) bool {
	val, _ := strconv.Atoi(string(runes))

	if val >= 1 && val <= 26 {
		return true
	}

	return false
}
