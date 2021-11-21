package X

func rotate(matrix [][]int) {
	l := len(matrix)
	xm, ym := 0, 0

	for n := l - 1; n > 0; n -= 2 {
		xmx, ymx := xm+n, ym+n

		for m := 0; m < n; m++ {
			x, y := xm, ym+m

			t := matrix[y][ymx]
			matrix[y][ymx] = matrix[x][y]

			t2 := matrix[xmx][ym+ymx-y]
			matrix[xmx][ym+ymx-y] = t

			t3 := matrix[xm+xmx-y][ym]
			matrix[xm+xmx-y][ym] = t2

			matrix[x][y] = t3
		}
		xm++
		ym++
	}
}
