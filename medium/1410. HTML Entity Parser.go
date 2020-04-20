package problem

func EntityParser(text string) string {
	var start int
	mapping := map[string]rune{
		"&quot":  '"',
		"&apos":  '\'',
		"&amp":   '&',
		"&gt":    '>',
		"&lt":    '<',
		"&frasl": '/',
	}

	for i := 0; i < len(text); i++ {
		t := text[i]
		if t == '&' {
			start = i
		}

		if t == ';' && start != -1 {
			matchString := text[start:i]
			val, ok := mapping[matchString]
			if ok {
				var suffix string
				if i != len(text) {
					suffix = text[i+1:]
				}
				text = text[:start] + string(val)
				text += suffix
				i = start
				start = -1
			}
		}
	}

	return text
}
