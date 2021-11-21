package X

func permuteUnique(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	res := make([][]int, 0)
	history := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if history[nums[i]]{
			continue
		}

		history[nums[i]] = true

		nums[0], nums[i] = nums[i], nums[0]
		tRes := permuteUnique(nums[1:])

		for j := 0; j < len(tRes); j++ {
			tRes[j] = append([]int{nums[0]}, tRes[j]...)
		}

		nums[0], nums[i] = nums[i], nums[0]
		res = append(res, tRes...)
	}

	return res
}
