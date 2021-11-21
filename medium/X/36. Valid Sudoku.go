package X

import "fmt"

func isValidSudoku(board [][]byte) bool {
	rowMap := map[int]map[byte]bool{}
	colMap := map[int]map[byte]bool{}
	gridMap := map[string]map[byte]bool{}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if !(board[i][j] >= '0' && board[i][j] <= '9') {
				continue
			}

			// check row
			if row, ok := rowMap[i]; ok {
				if _, ok := row[board[i][j]]; ok {
					return false
				}
				row[board[i][j]] = true
			} else {
				rowMap[i] = map[byte]bool{}
				rowMap[i][board[i][j]] = true
			}

			// check col
			if col, ok := colMap[j]; ok {
				if _, ok := col[board[i][j]]; ok {
					return false
				}
				col[board[i][j]] = true
			} else {
				colMap[j] = map[byte]bool{}
				colMap[j][board[i][j]] = true
			}

			// NICE: could have used key = (r / 3) * 3 + c / 3 instead
			// check grid
			gridKey := fmt.Sprintf("%v%v", i/3, j/3)
			if grid, ok := gridMap[gridKey]; ok {
				if _, ok := grid[board[i][j]]; ok {
					return false
				}
				grid[board[i][j]] = true
			} else {
				gridMap[gridKey] = map[byte]bool{}
				gridMap[gridKey][board[i][j]] = true
			}
		}
	}

	return true
}
