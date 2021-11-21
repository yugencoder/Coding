package X

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	a := binaryLeftSearch(nums, target)
	b := binaryRightSearch(nums, target)

	return []int{a, b}
}

func binaryLeftSearch(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := start + (end-start)/2
		if target <= nums[mid] {
			if nums[mid] == target && ((mid == 0) || (mid > 0 && nums[mid-1] != target)) {
				return mid
			}
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

func binaryRightSearch(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := start + (end-start)/2
		if target >= nums[mid] {
			if nums[mid] == target && ((mid == len(nums)-1) || (mid < len(nums) && nums[mid+1] != target)) {
				return mid
			} else {
				start = mid + 1
			}
		} else {
			end = mid - 1
		}
	}
	return -1
}