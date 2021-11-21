package easy

import "fmt"

func isPalindrome(x int) bool {
	palStr := fmt.Sprintf("%d", x)

	for i, j := 0, len(palStr)-1; i <= j; i, j = i+1, j-1 {
		if palStr[i] != palStr[j] {
			return false
		}
	}

	return true
}
