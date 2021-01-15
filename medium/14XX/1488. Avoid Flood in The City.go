package _14XX

func avoidFlood(rains []int) []int {
	arr := []int{}

	rainMap := map[int]int{}
	next := []int{}
	for _, r := range rains {
		if r == 0 {
			continue
		}

		if rainMap[r] > 0 {
			next = append(next, r)
		}
		rainMap[r] += 1
	}

	lakeCount := map[int]int{}
	for _, r := range rains {
		if r > 0 {
			if lakeCount[r] > 0 {
				return []int{}
			}
			arr = append(arr, -1)
		} else {
			if len(next) == 0 {
				arr = append(arr, 1)
				continue
			}

			hasNext := false
			// nothing to fill - no repeat yet
			for i, nxt := range next {
				if lakeCount[nxt] > 0 {
					hasNext = true
					arr = append(arr, nxt)

					lakeCount[nxt] -= 1
					if i == len(next) {
						next = next[:len(next)-1]
					} else {
						next = append(next[:i], next[i+1:]...)
					}
					break
				}
			}

			if !hasNext{
				arr = append(arr, 1)
			}

		}

		lakeCount[r] += 1
	}

	return arr
}
