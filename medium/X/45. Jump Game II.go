package X

import "math"

func jump(nums []int) int {
	currFarthest := 0
	nextFarthest := 0
	x := 0

	for i, n := range nums {
		if currFarthest >= len(nums)-1 {
			return x
		}

		if i+n > nextFarthest {
			nextFarthest = i + n
		}

		if i >= currFarthest {
			currFarthest = nextFarthest
			x++
		}
	}

	return 0
}

func jump2(nums []int) int {
	mem := make([]int, len(nums))
	for i := range mem {
		mem[i] = math.MaxInt32
	}

	mem[0] = 0
	for i, n := range nums {
		for j := 1; j <= n && i+j < len(nums); j++ {
			mem[i+j] = min(mem[i+j], mem[i]+1)
		}
	}

	return mem[len(nums)-1]
}
