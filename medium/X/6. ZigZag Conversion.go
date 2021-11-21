package X

func convert(s string, numRows int) string {
	res := []rune{}
	diag := max(0, numRows-2)

	for i := 0; i < numRows; i++ {
		curr := i
		down := 1

		for curr < len(s) {
			res = append(res, rune(s[curr]))

			if i == 0 || i == numRows-1 {
				curr += numRows + diag
			} else if down > 0 {
				curr += 2 * (numRows - 1 - i)
			} else {
				curr += 2 * i
			}
			down *= -1
		}
	}

	return string(res)
}
