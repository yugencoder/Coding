package X

var ans [][]int

func subsetsWithDup(nums []int) [][]int {
	countMap := map[int]int{}
	nonRepeatedNums := []int{}

	for _, n := range nums {
		if countMap[n] == 0 {
			nonRepeatedNums = append(nonRepeatedNums, n)
		}
		countMap[n]++
	}

	ans = [][]int{{}}

	for i := range nonRepeatedNums {
		findSubsets(i, nonRepeatedNums, []int{}, countMap)
	}

	return ans

}

func findSubsets(start int, nums []int, curr []int, countMap map[int]int) {
	if start >= len(nums) {
		return
	}

	if countMap[nums[start]]--; countMap[nums[start]] < 0 {
		countMap[nums[start]]++
		return
	}

	curr = append(curr, nums[start])
	ans = append(ans, append([]int{}, curr...))

	for i := start; i < len(nums); i++ {
		if countMap[nums[i]] >= 0 {
			findSubsets(i, nums, curr, countMap)
		} else {
			findSubsets(i+1, nums, curr, countMap)
		}
	}

	countMap[nums[start]]++

}

func getMapCopy(countMap map[int]int) map[int]int {
	nMap := make(map[int]int)
	for k, v := range countMap {
		nMap[k] = v
	}

	return nMap
}
