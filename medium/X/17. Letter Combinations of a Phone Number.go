package X

func letterCombinations(digits string) []string {
	if digits == ""{
		return []string{}
	}

	digitByChars := map[string][]rune{
		"2": {'a', 'b', 'c'},
		"3": {'d', 'e', 'f'},
		"4": {'g', 'h', 'i'},
		"5": {'j', 'k', 'l'},
		"6": {'m', 'n', 'o'},
		"7": {'p', 'q', 'r', 's'},
		"8": {'t', 'u', 'v'},
		"9": {'w', 'x', 'y', 'z'},
	}
	res := []string{""}
	for _, d := range digits {
		newRes := []string{}

		for _, v := range digitByChars[string(d)] {
			for _, r := range res {
				newRes = append(newRes, r+string(v))
			}
		}
		res = newRes
	}

	return res
}
