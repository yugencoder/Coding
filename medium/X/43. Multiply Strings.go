package X

import (
	"fmt"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0"{
		return "0"
	}

	nums := make([]int, len(num1)+len(num2)-1)
	offset := 0

	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			n1 := int(num1[i] - '0')
			n2 := int(num2[j] - '0')

			res := n1*n2
			nums[offset+len(num2)-1-j] += res
		}
		offset++
	}

	res := ""
	c := 0
	for i := 0; i < len(nums); i++ {
		nums[i] += c
		c = nums[i] / 10
		res = fmt.Sprintf("%v%s", nums[i]%10, res)
	}

	if c > 0 {
		res = fmt.Sprintf("%v%s", c, res)
	}

	return res
}
