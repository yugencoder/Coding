package X

var numMap map[int]int

func numTrees(n int) int {
	if n <= 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	res := 0
	if numMap == nil {
		numMap = map[int]int{}
	}

	if v := numMap[n]; v > 0 {
		return v
	}

	for i := 0; i < n; i++ {
		res += numTrees(i) * numTrees(n-1-i)
	}

	numMap[n] = res

	return res
}
