package X

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))

	for j := 0; j < len(board); j++ {
		visited[j] = make([]bool, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && dfs(i, j, 1, board, word, getCopy(visited)) {
				return true
			}
		}
	}

	return false
}

func dfs(i int, j int, k int, board [][]byte, word string, visited [][]bool) bool {
	if k == len(word) {
		return true
	}

	if visited[i][j] == true {
		return false
	}

	dir := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	visited[i][j] = true

	for _, d := range dir {
		nI, nJ := i+d[0], j+d[1]

		if nI < len(board) && nI >= 0 && nJ >= 0 && nJ < len(board[0]) {
			if board[nI][nJ] == word[k] {
				if dfs(nI, nJ, k+1, board, word, getCopy(visited)) {
					return true
				}
			}
		}

	}

	return false
}

func getCopy(matrix [][]bool) [][]bool {
	duplicate := make([][]bool, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]bool, len(matrix[0]))
		copy(duplicate[i], matrix[i])
	}

	return duplicate
}
