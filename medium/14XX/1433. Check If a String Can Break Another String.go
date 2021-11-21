package _14XX

import "sort"

func checkIfCanBreak(s1 string, s2 string) bool {
	var s1Arr, s2Arr []rune

	for i := 0; i < len(s1); i++ {
		s1Arr= append(s1Arr, rune(s1[i]))
		s2Arr =  append(s2Arr, rune(s2[i]))
	}

	sort.Slice(s1Arr, func(i, j int) bool {
		return s1Arr[i] < s1Arr[j]
	})

	sort.Slice(s2Arr, func(i, j int) bool {
		return s2Arr[i] < s2Arr[j]
	})

	leftOrder := false
	rightOrder := false
	for i := 0; i < len(s1); i++ {
		if s1Arr[i] < s2Arr[i] {
			leftOrder = true
		} else if s2Arr[i] < s1Arr[i] {
			rightOrder = true
		}
		if leftOrder && rightOrder {
			return false
		}
	}
	return true
}
