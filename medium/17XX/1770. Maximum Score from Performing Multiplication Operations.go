package _7XX


func maximumScore2(nums []int, multipliers []int) int {
	mem := make([][]int, len(multipliers))

	for n := range mem {
		mem[n] = make([]int, len(multipliers))
		for i := range mem[n] {
			mem[n][i] = -1000000001
		}
	}

	return helper(nums, multipliers, 0, len(nums)-1, 0,mem)
}

func helper(nums, mul []int, start, end int, i int,mem [][]int) int {
	if i == len(mul) {
		return 0
	}

	if mem[start][i] > -1000000001 {
		return mem[start][i]
	}

	left := mul[i]*nums[start] + helper(nums, mul, start+1, end, i+1,mem)
	right := mul[i]*nums[end] + helper(nums, mul, start, end-1, i+1,mem)

	mem[start][i] = max(left, right)

	return mem[start][i]
}
