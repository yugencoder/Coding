package _5XX

func numSubmat(mat [][]int) int {
	res := 0

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 {
				// check all combinations
				for m := 1; m <= len(mat)-i; m++ {
					for n := 1; n <= len(mat[0])-j; n++ {
						allTrue := true

						for r := 0; r < m; r++ {
							for s := 0; s < n; s++ {
								if mat[i+r][j+s] == 0 {
									allTrue = false
									break
								}
							}
							if !allTrue {
								break
							}
						}
						if allTrue {
							res++
						}
					}
				}
			}
		}
	}

	return res
}
