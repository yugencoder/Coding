package easy

func RomanToInt(s string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	romanSpecialMap := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	var skip bool
	res := 0

	for i := range s {
		if skip {
			skip = false
			continue
		}

		if i+1 < len(s) {
			if v, ok := romanSpecialMap[s[i:i+2]]; ok {
				res += v
				skip = true
				continue
			}
		}

		if v, ok := romanMap[s[i:i+1]]; ok {
			res += v
		}
	}

	return res
}
