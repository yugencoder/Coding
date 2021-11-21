package X

func removeDuplicates(nums []int) int {
	w := 0
	limit := 1
	match := 0

	var p *int

	for _, n := range nums {
		if p != nil && n == *p {
			match++
		} else {
			match = 0
		}

		if match <= limit || p == nil {
			nums[w] = n
			w++
		}

		p = getIntCopy(n)

	}

	return w
}

func getIntCopy(n int) *int {
	a := n

	return &a
}
