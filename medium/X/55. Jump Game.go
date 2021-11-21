package X

func canJump(nums []int) bool {
	maxVal := 0

	for i, n := range nums {
		if i <= maxVal {
			maxVal = max(maxVal, i+n)
		}
	}

	if maxVal >= len(nums)-1 {
		return true
	}

	return false
}
