package t13XX

func xorQueries(arr []int, queries [][]int) []int {
	var xor, ans []int
	var xorSum int

	xor = make([]int, len(arr)+1)
	ans = make([]int, len(queries))

	xor[0] = 0
	for i, a := range arr {
		xorSum ^= a
		xor[i+1] = xorSum
	}

	for i, q := range queries {
		s1 := xor[q[0]]
		s2 :=  xor[q[1]+1]

		ans[i] = s1 ^ s2
	}

	return ans
}
