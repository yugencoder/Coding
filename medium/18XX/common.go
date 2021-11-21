package _8XX

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func getGreaterThanEqualTo(n int, nums [][]int, start, end int) int {
	mid := start + (end-start)/2

	if nums[mid][0] == n || end == start {
		return nums[mid][1]
	}

	if nums[mid][0] <= n {
		if mid == len(nums)-1 {
			return nums[mid][1]
		}
		return getGreaterThanEqualTo(n, nums, mid+1, end)
	} else {
		return getGreaterThanEqualTo(n, nums, start, mid)
	}
}

func getLesserThanEqualTo(n int, nums [][]int, start, end int) int {
	mid := start + (end-start+1)/2

	if nums[mid][0] == n || end == start {
		return nums[mid][1]
	}

	if nums[mid][0] >= n {
		if mid == 0 {
			return nums[mid][1]
		}
		return getLesserThanEqualTo(n, nums, start, mid-1)
	} else {
		return getLesserThanEqualTo(n, nums, mid, end)
	}
}
