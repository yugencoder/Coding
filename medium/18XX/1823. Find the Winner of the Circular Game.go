package _8XX

func findTheWinner(n int, k int) int {
	circle := make([]int, n)
	k = k - 1

	for i := 0; i < n; i++ {
		circle[i] = i + 1
	}

	for i := 0; len(circle) > 1; {
		i, circle = removeCircle(circle, i, k)
	}

	return circle[0]
}

func removeCircle(circle []int, i, k int) (int, []int) {
	newI := (i + k) % len(circle)

	if newI < len(circle)-1 {
		circle = append(circle[:newI], circle[newI+1:]...)
	} else {
		circle = circle[:newI]
	}

	return newI % len(circle), circle
}
