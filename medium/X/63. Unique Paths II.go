package X

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	res := make([][]int, m)

	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		res[i][0] = 1
	}

	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}

		res[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}

			if obstacleGrid[i-1][j] == 0 {
				res[i][j] += res[i-1][j]
			}

			if obstacleGrid[i][j-1] == 0 {
				res[i][j] += res[i][j-1]
			}
		}
	}

	return res[m-1][n-1]
}
