package _7XX

func minNoOperations(boxes string) []int {
	res := []int{}

	left := []int{}
	right := []int{}

	for i, b := range boxes {
		if b == '1' && i > 0 {
			right = append(right, i)
		}
	}

	for _, b := range boxes {
		lSum := 0
		rSum := 0

		newLeft := []int{}
		for _, l := range left {
			lSum += l
			newLeft = append(newLeft, l+1)
		}

		if b == '1' {
			newLeft = append(newLeft, 1)
		}

		newRight := []int{}
		for _, r := range right {
			rSum += r
			if r == 1 {
				continue
			}
			newRight = append(newRight, r-1)
		}

		left = newLeft
		right = newRight
		res = append(res, lSum+rSum)
	}

	return res
}
