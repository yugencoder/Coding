package _XX

func nextGreaterElements(nums []int) []int {
	stack := make([]int, 0)
	res := make([]int, len(nums))

	for i := 0; i < len(res); i++ {
		res[i] = -1
	}

	for i := len(nums)*2 - 1; i >= 0; i-- {
		j := i % len(nums)
		curr := nums[j]

		for ; len(stack) > 0 && curr >= stack[len(stack)-1]; {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 && curr < stack[len(stack)-1] {
			res[j] = stack[len(stack)-1]
		}

		// push to stack
		for ; len(stack) > 0 && nums[j] >= stack[len(stack)-1]; {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, curr)

	}

	return res
}
