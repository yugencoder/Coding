package easy

import "fmt"

func FindMissingRanges(nums []int, lower int, upper int) []string {
	nums = append([]int{lower-1}, nums...)
	nums = append(nums, upper+1)

	res := make([]string,0)

	for i, n := range nums {
		f := n
		s := n

		if i+1 <len(nums){
			s = nums[i+1]
		}

		if f +1 < s {
			res = append(res, printVal(f+1, s-1))
		}
	}

	return res
}

func printVal(a, b int) string {
	if a == b {
		return fmt.Sprintf("%v", a)
	}
	return fmt.Sprintf("%v->%v", a, b)
}
