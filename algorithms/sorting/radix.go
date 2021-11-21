package sorting

func maxBit(arr []int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	bit := 1
	for max >= 10 {
		max /= 10
		bit++
	}
	return bit
}

func RadixSort(arr []int) {
	mBit := maxBit(arr)
	tmp := make([]int, len(arr), len(arr))
	count := make([]int, 10, 10) // Counter
	for i, radix := 1, 1; i <= mBit; i++ {
		// Clear the counter before each allocation
		for j := 0; j < 10; j++ {
			count[j] = 0
		}

		for j := 0; j < len(arr); j++ {
			k := (arr[j] / radix) % 10 // Count the number of records in each bucket
			count[k]++
		}

		for j := 1; j < 10; j++ {
			count[j] = count[j-1] + count[j] // Assign the position in tmp to each bucket in turn
		}

		// Note: Descending order
		for j := len(arr) - 1; j >= 0; j-- {
			// Collect records in all buckets into tmp in turn
			k := (arr[j] / radix) % 10
			tmp[count[k]-1] = arr[j]
			count[k]--
		}

		copy(arr, tmp)
		radix *= 10
	}
}
