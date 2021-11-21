package _8XX

func minSideJumps2(obstacles []int) int {
	curr := []int{1, 0, 1}

	for _, o := range obstacles {
		if o > 0 {
			curr[o-1] = 10000000
		}

		for j := 0; j < 3; j++ {
			if j != o-1 {
				curr[j] = min(curr[j], min(1+curr[(j+1)%3], 1+curr[(j+2)%3]))
			}
		}
	}

	return min(min(curr[0], curr[1]), curr[2])
}
