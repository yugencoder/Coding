package aiq

type prevData struct {
	idx   int
	count int
}

func distanceMetric(nums []int) []int {
	res := make([]int, len(nums))
	left := make([]int, len(nums))
	right := make([]int, len(nums))

	lastMap := map[int]prevData{}

	for i, n := range nums {
		prev, ok := lastMap[n]
		if !ok {
			prev = prevData{
				count: 0,
			}

		} else {
			left[i] = left[prev.idx] + (i-prev.idx)*prev.count
		}

		lastMap[n] = prevData{
			idx:   i,
			count: prev.count + 1,
		}

	}

	lastMap = map[int]prevData{}
	for i := len(nums) - 1; i >= 0; i-- {
		n := nums[i]
		prev, ok := lastMap[n]
		if !ok {
			prev = prevData{
				count: 0,
			}

		} else {
			right[i] = right[prev.idx] + (prev.idx-i)*prev.count
		}

		lastMap[n] = prevData{
			idx:   i,
			count: prev.count + 1,
		}
	}

	for i := range nums {
		res[i] = left[i] + right[i]
	}

	return res
}

func distanceMetric2(nums []int) []int {
	res := make([]int, len(nums))

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				res[i] += j - i
				res[j] += j - i
			}
		}
	}

	return res
}
