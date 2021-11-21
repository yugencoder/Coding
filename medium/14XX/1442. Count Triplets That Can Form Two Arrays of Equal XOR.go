package _14XX

func countTriplets(arr []int) int {
	res := 0
	type valWithIndex struct {
		val   int
		index int
	}
	next := []valWithIndex{}
	curr := []valWithIndex{}

	for i, a := range arr {
		if len(curr) == 0 {
			curr = append(curr, valWithIndex{a,i})
			continue
		}

		next = []valWithIndex{}
		for _, c := range curr {
			val := c.val ^ a
			if val == 0 {
				res += i - c.index
			}
			next = append(next, valWithIndex{
				val:   val,
				index: c.index,
			})
		}

		next = append(next, valWithIndex{a,i})
		curr = next
	}

	return res
}
