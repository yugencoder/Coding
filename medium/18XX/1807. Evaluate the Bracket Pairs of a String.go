package _8XX

func evaluate(s string, knowledge [][]string) string {
	var ans string
	start := -1
	end := 0
	kMap := map[string]string{}

	for _, k := range knowledge {
		kMap[k[0]] = k[1]
	}

	for i, c := range s {
		if c == '(' {
			start = i + 1
		} else if c == ')' && start != -1 {
			if start-1 >= end {
				ans += s[end : start-1]
			}

			if v, ok := kMap[s[start:i]]; ok {
				ans += v
			} else {
				ans += "?"
			}

			start = -1
			end = i + 1
		}
	}

	if end < len(s) {
		return ans + s[end:]
	}

	return ans
}
