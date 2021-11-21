package easy

func RemoveDuplicates(nums []int) int {
	k := 0
	for i, n := range nums{
		if i == 0{
			continue
		}

		if n != nums[k]{
			nums[k+1] = n
			k++
		}
	}

	return k+1
}
