package _14XX

import "container/list"

func longestSubarray(nums []int, limit int) int {
	var res int

	l := 0
	minList := list.New()
	maxList := list.New()

	for r, n := range nums {
		for ; minList.Len() > 0 && n < minList.Back().Value.(int); {
			minList.Remove(minList.Back())
		}

		for ; maxList.Len() > 0 && n > maxList.Back().Value.(int); {
			maxList.Remove(maxList.Back())
		}

		minList.PushBack(n)
		maxList.PushBack(n)

		if maxList.Front().Value.(int)-minList.Front().Value.(int) > limit {
			if maxList.Front().Value.(int) == nums[l] {
				maxList.Remove(maxList.Front())
			}

			if minList.Front().Value.(int) == nums[l] {
				minList.Remove(minList.Front())
			}

			l++
		}

		res = max(res, r-l+1)
	}

	return res
}
