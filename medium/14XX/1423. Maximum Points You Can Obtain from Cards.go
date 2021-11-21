package _14XX

func maxScore(cardPoints []int, k int) int {
	first := cardPoints[0:k]
	second := cardPoints[len(cardPoints)-k:]
	second = reverse(second)
	firstSum, secondSum := make([]int, k), make([]int, k)

	sumF, sumS := 0, 0
	for i, f := range first {
		sumF += f
		sumS += second[i]

		firstSum[i] = sumF
		secondSum[i] = sumS
	}

	res := 0
	for i := 0; i < k; i++ {
		partialSum := firstSum[i]
		if k-i-2 >= 0 {
			partialSum += secondSum[k-i-2]
		}
		res = max(res, partialSum)

		partialSum = secondSum[i]
		if k-i-2 >= 0 {
			partialSum += firstSum[k-i-2]
		}
		res = max(res, partialSum)
	}

	return res
}

func reverse(arr []int) []int {
	copy := make([]int, len(arr))
	for i := 0; i < len(copy); i++ {
		copy[i] = arr[i]
	}

	j := len(copy) - 1
	for i := 0; i < len(copy)/2; i++ {
		copy[i], copy[j] = copy[j], copy[i]
		j--
	}

	return copy
}
