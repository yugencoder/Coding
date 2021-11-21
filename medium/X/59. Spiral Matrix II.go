package X

func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}

	res := make([][]int, n)
	for k := 0; k < n; k++ {
		res[k] = make([]int, n)
	}

	x := 1
	i, j := 0, 0
	coord := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	ic := 0
	next := coord[ic]

	//- shape
	for k := 0; k < n; k++ {
		res[i][j] = x
		i, j = i+next[0], j+next[1]
		x++
	}

	i, j = i-next[0], j-next[1]
	n--

	// L shape
	for n > 0 {
		ic++
		next = coord[ic%(len(coord))]
		for k := 0; k < n; k++ {
			i, j = i+next[0], j+next[1]
			res[i][j] = x
			x++
		}

		ic++
		next = coord[ic%(len(coord))]
		for k := 0; k < n; k++ {
			i, j = i+next[0], j+next[1]
			res[i][j] = x
			x++
		}

		n--
	}

	return res
}
