package X

func sortColors(nums []int) {
	r := 0
	i := 0
	j := len(nums) - 1

	for i <= j {
		if nums[i] == r {
			i++
			continue
		}
		if nums[j] != r {
			j--
			continue
		}

		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	r = 1
	j = len(nums) - 1
	for i < j {
		if nums[i] == r {
			i++
			continue
		}
		if nums[j] != r {
			j--
			continue
		}

		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
