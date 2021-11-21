package X

func searchMatrix(matrix [][]int, target int) bool {
	r := binaryColSearch(matrix, 0, target, 0, len(matrix)-1)
	c := binaryRowSearch(matrix[r], target, 0, len(matrix[0])-1)

	return matrix[r][c] == target

}

func binaryColSearch(matrix [][]int, i int, target int, start, end int) int {
	if start <= end {
		mid := start + (end-start)/2

		if matrix[mid][i] == target || start == end {
			return mid
		}

		if target < matrix[mid][i] {
			return binaryColSearch(matrix, i, target, start, mid-1)
		} else {
			if mid+1 == end {
				if matrix[end][i] <= target {
					return end
				} else {
					return mid
				}
			}

			return binaryColSearch(matrix, i, target, mid, end)
		}
	}

	return start
}

func binaryRowSearch(col []int, target int, start, end int) int {
	if start <= end {
		mid := start + (end-start)/2

		if col[mid] == target || start == end {
			return mid
		}

		if col[mid] > target {
			return binaryRowSearch(col, target, start, mid-1)
		} else {
			if mid+1 == end {
				if col[end] <= target {
					return end
				} else {
					return mid
				}
			}

			return binaryRowSearch(col, target, mid, end)
		}
	}

	return start
}
