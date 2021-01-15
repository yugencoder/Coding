package t13XX

import (
	"fmt"
	"testing"
)

func TestValidateBinaryTreeNodes(t *testing.T) {
	//n := 4
	//left := []int{1, -1, 3, -1}
	//right := []int{2, -1, -1, -1}

	//5
	//[1,3,-1,-1,-1]
	//[-1,2,4,-1,-1]

	n := 5
	left := []int{1,3,-1,-1,-1}
	right := []int{-1,2,4,-1,-1}

	res := validateBinaryTreeNodes(n, left, right)
	fmt.Print(res)
}
