package _8XX

func reinitializePermutation(n int) int {
	count := 0
	if n == 2 {
		return 1
	}

	i := 1
	init := 1
	for ; ; {
		if i%2 == 1 {
			i = n/2 + (i-1)/2
		} else {
			i = i / 2
		}
		count++
		if i == init {
			break
		}

	}

	return count
}
