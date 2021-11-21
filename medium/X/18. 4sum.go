package X

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			localTarget := target - nums[i] - nums[j]
			start := j + 1
			end := len(nums) - 1

			for start < end {
				currSum := nums[start] + nums[end]
				if currSum == localTarget {
					res = append(res, []int{nums[i], nums[j], nums[start], nums[end]})

					for start < end && nums[start] == nums[start+1] {
						start++
					}
					for start < end && nums[end] == nums[end-1] {
						end--
					}

					start++
					end--
				} else if currSum < localTarget {
					start++
				} else {
					end--
				}
			}
		}
	}

	return res
}
