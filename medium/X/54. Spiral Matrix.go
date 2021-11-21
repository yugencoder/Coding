package X

func spiralOrder(matrix [][]int) []int {
	res := []int{}
	r := len(matrix)
	c := len(matrix[0])

	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	d := 0
	i, j := 0, -1
	isRow := true
	for r > 0 && c > 0 {
		x, y := dir[d][0], dir[d][1]
		l := r
		if isRow {
			l = c
		}

		for m := l; m > 0; m-- {
			i, j = i+x, j+y
			res = append(res, matrix[i][j])

		}

		if isRow {
			r--
		} else {
			c--
		}
		d++
		d %= 4
		isRow = !isRow
	}

	return res
}
