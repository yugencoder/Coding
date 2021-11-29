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
		if prev, ok := lastMap[n]; !ok {
			lastMap[n] = prevData{
				idx:   i + 1,
				count: 1,
			}
			left[i] = (prev.idx)*prev.count + i + 1 - prev.idx
		}

		prev := lastMap[n]
		prev.count += 1
		prev.idx = i
	}

	lastMap = map[int]prevData{}
	for i := len(nums) - 1; i >= 0; i-- {
		n := nums[i]
		if prev, ok := lastMap[n]; !ok {
			lastMap[n] = prevData{
				idx:   i + 1,
				count: 1,
			}
			right[i] = (len(nums)-prev.idx)*prev.count + prev.idx + 1 - i
		}

		prev := lastMap[n]
		prev.count += 1
		prev.idx = i
	}

	for i := range nums {
		res[i] = left[i] + right[i]
	}

	return res
}
