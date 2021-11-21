package _8XX

func maxValue(n int, index int, maxSum int) int {

	for avg := maxSum / n; ; avg++ {
		sum := avg

		lt := index
		rt := n - 1 - index

		avg--
		if avg == 0 {
			sum = index
		} else {
			if avg > lt {
				sum += avg*(avg+1)/2 - ((avg - lt) * (avg - lt + 1) / 2)
			} else {
				sum += avg*(avg+1)/2 + lt - avg
			}

			if avg > rt {
				sum += avg*(avg+1)/2 - ((avg - rt) * ((avg - rt) + 1) / 2)
			} else {
				sum += avg*(avg+1)/2 + rt - avg
			}
		}

		if maxSum < sum {
			return avg
		}

		avg++
	}

	return 0
}
