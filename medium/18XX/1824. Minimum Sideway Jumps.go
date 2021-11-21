package _8XX

func minSideJumps(obstacles []int) int {
	var sideJumps, curr int
	curr = 2

	nextNextIndex := 0
	nextNext := 0

	for i := range obstacles {
		next := 0

		if i+1 < len(obstacles) {
			next = obstacles[i+1]
		}

		if nextNextIndex < i+2 {
			nextNextIndex, nextNext = nextNextObstacle(obstacles, next, i+2)
		}

		if next != 0 && curr == next {
			if nextNext == 0 {
				curr = 5 //fictional
			} else {
				curr = 1 ^ 2 ^ 3 ^ next ^ nextNext
				if obstacles[i] == curr {
					curr = nextNext
				}
			}

			sideJumps++
		}

	}

	return sideJumps
}

func nextNextObstacle(obstacles []int, next int, startIndex int) (int, int) {
	for i := startIndex; i < len(obstacles); i++ {
		if obstacles[i] > 0 && next != obstacles[i] {
			return i, obstacles[i]
		}
	}

	return 0, 0
}
