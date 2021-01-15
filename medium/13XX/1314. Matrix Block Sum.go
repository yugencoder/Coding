package t13XX

import "fmt"

func matrixBlockSum(mat [][]int, K int) [][]int {
	res := make([][]int, len(mat))

	for i := range mat {
		res[i] = make([]int, len(mat[i]))
	}

	sumArr := createSumArr(mat)
	fmt.Println(sumArr)

	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[0]); j++ {
			iK := len(res) - 1
			jK := len(res[0]) - 1

			if i+K < len(res) {
				iK = i + K
			}

			if j+K < len(res[0]) {
				jK = j + K
			}

			mik := 0
			mjk := 0
			mijk := 0

			if i-1-K >= 0 {
				mik = sumArr[i-1-K][jK]
			}
			if j-1-K >= 0 {
				mjk = sumArr[iK][j-1-K]
			}

			if i-1-K >= 0 && j-1-K >= 0 {
				mijk = sumArr[i-1-K][j-1-K]
			}


			res[i][j] = sumArr[iK][jK] - (mik + mjk - mijk)
		}
	}

	return res
}

func createSumArr(mat [][]int) [][]int {
	res := make([][]int, len(mat))

	for i := range mat {
		res[i] = make([]int, len(mat[i]))
	}

	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[0]); j++ {
			ans := 0
			if i > 0 {
				ans += res[i-1][j]
			}

			if j > 0 {
				ans += res[i][j-1]
			}

			if i > 0 && j > 0 {
				ans -= res[i-1][j-1]
			}

			res[i][j] = ans + mat[i][j]
		}
	}

	return res
}
