package X

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	res := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		nums[0], nums[i] = nums[i], nums[0]
		tRes := permute(nums[1:])

		for j := 0; j < len(tRes); j++ {
			tRes[j] = append([]int{nums[0]}, tRes[j]...)
		}

		nums[0], nums[i] = nums[i], nums[0]
		res = append(res, tRes...)
	}

	return res
}
