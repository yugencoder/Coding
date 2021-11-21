package X

import (
	"fmt"
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var res []string

	getAllAddresses(0, []string{}, &res, s)

	fmt.Println(res)
	return res
}

func getAllAddresses(start int, curr []string, res *[]string, s string) {
	if start == len(s) && len(curr) == 4 {
		*res = append(*res, strings.Join(curr, "."))

		return
	}

	if len(s)-start > (4-len(curr))*3 || len(curr) > 3 {
		return
	}

	for i := 1; i <= 3; i++ {
		if start+i <= len(s) && !(string(s[start]) == "0" && i > 1) {
			v, err := strconv.Atoi(s[start : start+i])
			if err == nil && v <= 255 {
				getAllAddresses(start+i, append(curr, s[start:start+i]), res, s)
			}
		}
	}
}
