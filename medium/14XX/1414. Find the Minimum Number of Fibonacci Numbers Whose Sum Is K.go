package _14XX

func findMinFibonacciNumbers(k int) int {
	res := 0

	nos := findMinFibonacciNumbersLessThan(k)

	for ; ; {
		n := findNoLessThanK(nos, 0, len(nos)-1, k)
		n = nos[n]
		if n == k {
			return res + 1
		} else {
			res += k/n
			k = k % n
		}
	}
	return 0
}

func findNoLessThanK(array []int, start_idx, end_idx, search_val int) int {
	if start_idx == end_idx {
		if array[start_idx] <= search_val {
			return start_idx
		} else {
			return -1
		}
	}

	mid_idx := start_idx + (end_idx-start_idx)/2

	if search_val < array[mid_idx] {
		return findNoLessThanK(array, start_idx, mid_idx, search_val)
	}

	ret := findNoLessThanK(array, mid_idx+1, end_idx, search_val)
	if ret == -1 {
		return mid_idx
	}

	return ret
}

func findMinFibonacciNumbersLessThan(k int) ([]int) {
	res := []int{1, 1}
	first := 1
	second := 1

	if k == first {
		return res
	}

	if k < 2 {
		return res
	}

	for ; second < k; {
		third := first + second
		first, second = second, third
		res = append(res, third)

	}

	return res
}
