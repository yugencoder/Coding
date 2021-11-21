package X

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	for i := 1; i < m; i++ {
		grid[i][0] = grid[i-1][0] + grid[i][0]
	}

	for i := 1; i < n; i++ {
		grid[0][i] = grid[0][i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = grid[i][j] + min(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[m-1][n-1]
}
