package easy

import (
	"fmt"
)

func findLUSlength(a string, b string) int {
	if len(b) < len(a) {
		a, b = b, a
	}

	i, j := 0, 0
	res := 0
	common := 0
	for ; i < len(a) && j < len(b); i = i + 1 {
		if a[i] == b[j] {
			j++
			res++
			common++
		} else {
			for ; j < len(b) && a[i] != b[j]; j++ {
				res++
			}

			if j < len(b) {
				j++
				common++
			}
		}
	}

	fmt.Println(res, common)
	if common == len(a) {
		return -1
	}

	return res
}
