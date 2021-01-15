package _14XX

func kthFactor(n int, k int) int {
	var res int

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			res++
			if res == k {
				return i
			}
		}
	}

	return -1
}
