package _7XX

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

