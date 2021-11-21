package _14XX

func findDiagonalOrder(nums [][]int) []int {
	n, res := len(nums), []int{}

	l := n
	j, maxRowLen := 0, 0

	for i := 0; i < l; i++ {
		_i := min(i, n-1)
		maxRowLen = max(maxRowLen, len(nums[_i]))
		l = max(l, n+maxRowLen-1)

		for t, u := _i, j; t >= 0; t, u = t-1, u+1 {
			if len(nums[t]) > u {
				res = append(res, nums[t][u])
			}

			if u == maxRowLen {
				break
			}
		}

		if _i == n-1 {
			j++
		}
	}

	return res
}
