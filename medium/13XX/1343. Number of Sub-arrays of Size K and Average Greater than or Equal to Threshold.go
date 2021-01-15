package t13XX

func numOfSubarrays(arr []int, k int, threshold int) int {
	res := 0
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]

		if i >= k {
			sum -= arr[i-k]
		}

		if sum/k >= threshold {
			res++
		}
	}

	return res
}
