package _14XX

func kLengthApart(nums []int, k int) bool {
	zeroL := k
	for _, n := range nums {
		if n == 1 {
			if zeroL < k {
				return false
			}
			zeroL = 0
			continue
		}
		zeroL++
	}

	return true
}
