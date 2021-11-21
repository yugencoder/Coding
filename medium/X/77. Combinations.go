package X

var res [][]int

func combine(n int, k int) [][]int {
	res = [][]int{}

	for i := 1; i <= n; i++ {
		backtrack(i+1, n, k, []int{i})
	}

	return res
}

func backtrack(start int, max int, size int, curr []int) {
	if len(curr) == size {
		temp := make([]int, len(curr))
		copy(temp, curr)

		res = append(res, temp)

		return
	}

	for l := start; l <= max; l++ {
		backtrack(l+1, max, size, append(curr, l))
	}
}
