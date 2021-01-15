package _14XX

func canArrange(arr []int, k int) bool {
	arrMap := map[int]int{}

	for _, a := range arr {
		/*
			Theory:
				(a+b)%k = 0
				a%k == x , b%k = k-x
		*/
		x := a % k
		if x < 0 {
			x += k
		}
		arrMap[x] += 1
	}

	if arrMap[0]%2 != 0 {
		return false
	}

	for i := 1; i <= k/2; i++ {
		if arrMap[i] != arrMap[k-i] {
			return false
		}
	}

	return true
}
