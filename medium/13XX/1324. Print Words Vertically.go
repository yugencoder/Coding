package t13XX

import (
	"fmt"
	"strings"
)

func printVertically(s string) []string {
	arr := strings.Split(s, " ")
	maxLen := 0
	res := []string{}

	for _, a := range arr {
		strings.TrimSpace(a)
		maxLen = max(maxLen, len(a))
	}

	for i := 0; i < maxLen; i++ {
		temp := ""
		for j := 0; j < len(arr); j++ {
			extraChar := " "
			if i < len(arr[j]) {
				extraChar = fmt.Sprintf("%v", string(arr[j][i]))
			}
			temp = fmt.Sprintf("%v%v", temp, extraChar)
		}
		temp = strings.TrimRight(temp, " ")
		if len(temp) > 0 {
			res = append(res, temp)
		}
	}

	return res
}
