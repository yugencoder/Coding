package X

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start >= 0 && end >= 0 {
		mid := start + (end-start)/2

		if nums[mid] == target {
			return mid
		} else if start >= end {
			return -1
		}

		if nums[start] <= nums[mid] {
			if target >= nums[start] && target <= nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return -1
}
