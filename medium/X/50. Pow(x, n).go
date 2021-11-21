package X

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = n * -1
	}

	if n == 1 {
		return x
	} else if n == 0 {
		return 1
	}

	ans := x
	next := x
	i := 1

	for ; i < n; i = 2 * i {
		ans = next
		next = ans * ans
	}

	return ans * myPow(x, n-(i/2))
}
