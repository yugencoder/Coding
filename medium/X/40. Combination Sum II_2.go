package X

func combinationSum2(nums []int, target int) (result [][]int) {
	visited := make([]int, 50)
	result = make([][]int, 0)
	for _, n := range nums {
		visited[n]++
	}
	newNums := make([]int, 0)
	for i, v := range visited {
		if v > 0 {
			newNums = append(newNums, i)
		}
	}

	combinationSum2Helper(0, target, []int{}, append([]int{}, visited...), newNums, &result)
	return result
}

func combinationSum2Helper(curr, target int, currRes, visited, nums []int, result *[][]int) {
	if target == 0 {
		*result = append(*result, append([]int{}, currRes...))

		return
	} else if target < 0 {
		return
	}

	for i := curr; i < len(nums); i++ {
		if visited[nums[i]] > 0 {
			visited[nums[i]]--
			//fmt.Println(i, target-nums[i], append(currRes, nums[i]), append([]int{}, visited...), nums, result)
			combinationSum2Helper(i, target-nums[i], append(currRes, nums[i]), append([]int{}, visited...), nums, result)
			visited[nums[i]]++
		}
	}
}
