package _XX


func nextGreaterElements(nums []int) []int {
	stack := make([]int, 0)
	res := make([]int, len(nums))
	//
	for i := 0; i < len(nums)*2; i++ {
		// pop from stack
		if (len(stack) > 0 && stack[len(stack)-1] > nums[i%len(nums)]) || len(stack) == 0 {
			stack = append(stack, nums[i%len(nums)])
			continue
		}

		for len(stack) > 0 && nums[(stack[len(stack)-1])%len(nums)] >= nums[i%len(nums)] {
			nums[i%len(nums)] = stack[0]
			stack = stack[1:]
		}
	}

	return res
}
