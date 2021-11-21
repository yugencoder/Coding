package X

func search2(nums []int, target int) bool {
	return isTargetAvailable(nums, 0, len(nums)-1, target)
}

func isTargetAvailable(nums []int, start, end int, target int) bool {
	mid := start + (end-start)/2

	if start > end {
		return false
	}

	if nums[mid] == target {
		return true
	} else if start == end {
		return false
	}

	if start+1 == end {
		return nums[start] == target || nums[end] == target
	}

	if nums[start] == nums[end] {
		return isTargetAvailable(nums, start, mid, target) || isTargetAvailable(nums, mid, end, target)
	}

	if target >= nums[mid] {
		if nums[start] > nums[mid] && target >= nums[start] {
			//go left
			return isTargetAvailable(nums, start, mid, target)
		}
		return isTargetAvailable(nums, mid, end, target)
		// go right
	} else {
		if nums[end] < nums[mid] && target <= nums[end] {
			//go right
			return isTargetAvailable(nums, mid+1, end, target)
		}
		return isTargetAvailable(nums, start, mid-1, target)
		//go left

	}

}
