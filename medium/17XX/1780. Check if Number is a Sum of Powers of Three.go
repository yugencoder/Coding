package _7XX

func checkPowersOfThree(n int) bool {
	if n == 0 {
		return true
	}

	if n%3 == 0 {
		return checkPowersOfThree(n / 3)
	}
	if n%3 == 1 {
		return checkPowersOfThree(n - 1)
	}

	return false
}
