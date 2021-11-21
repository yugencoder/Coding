package X

import "fmt"

func countAndSay(n int) string {
	res := "1"
	if n == 1 {
		return res
	}

	for i := 2; i <= n; i++ {
		count := 0
		newRes := ""
		if len(res) == 1 {
			newRes = fmt.Sprintf("%v%v", 1, res)
		} else {
			for j := 0; j < len(res)-1; j++ {
				count++
				if res[j] != []byte(res)[j+1] {
					newRes += fmt.Sprintf("%v%v", count, string(res[j]))
					count = 0
				}

				if j == len(res)-2 {
					if count == 0 {
						newRes += fmt.Sprintf("%v%v", 1, string(res[j+1]))
					} else {
						newRes += fmt.Sprintf("%v%v", count+1, string(res[j]))
					}
				}
			}
		}
		res = newRes
	}

	return res
}
