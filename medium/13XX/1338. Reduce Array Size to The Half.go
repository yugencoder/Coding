package t13XX

import "sort"

func minSetSize(arr []int) int {
	var countArr []int
	var res int

	sort.Ints(arr)
	freq := 0

	for i := 1; i < len(arr); i++ {
		if arr[i-1] == arr[i] {
			if i == len(arr)-1 {
				countArr = append(countArr, freq)
			}
			freq++
		} else {
			countArr = append(countArr, freq)
			freq = 1
		}
	}
	sort.Slice(countArr, func(i, j int) bool {
		return countArr[i] > countArr[j]
	})

	count := 0
	for _, c := range countArr {
		res++
		count += c
		if count >= len(arr)/2 {
			return res
		}
	}

	return res
}
