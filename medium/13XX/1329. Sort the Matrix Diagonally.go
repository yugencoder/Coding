package t13XX

import "sort"

func diagonalSort(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])

	j := 1
	sorted := make([][]int, m+n-1)
	for i := 0; i < m+n-1; i++ {
		sorted[i] = make([]int, j)
		if (m + n - 1 - i) > min(m, n) {
			j++
		} else {
			j--
		}
		if j > min(m, n) {
			j = min(m, n)
		}
	}

	//traverse and store
	rowIdx := 0
	colIdx := n - 1
	for i := 0; i < m+n-1; i++ {
		r := rowIdx
		c := colIdx

		for idx := 0; idx < len(sorted[i]); idx++ {
			sorted[i][idx] = mat[r][c]
			c++
			r++
		}

		colIdx--
		if colIdx < 0 {
			colIdx = 0
			rowIdx++
		}
	}

	for i := 0; i < len(sorted); i++ {
		sort.Ints(sorted[i])
	}

	//unpack
	rowIdx = 0
	colIdx = n - 1
	for i := 0; i < m+n-1; i++ {
		r := rowIdx
		c := colIdx

		for idx := 0; idx < len(sorted[i]); idx++ {
			mat[r][c] = sorted[i][idx]
			c++
			r++
		}

		colIdx--
		if colIdx < 0 {
			colIdx = 0
			rowIdx++
		}
	}

	return mat
}
