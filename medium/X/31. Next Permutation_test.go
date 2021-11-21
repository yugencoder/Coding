package X

import (
	"fmt"
	"testing"
)

func Test_nextPermutation(t *testing.T) {

	tests := []struct {
		name string
		nums []int
	}{
		{
			name: "",
			nums: []int{1, 5, 1, 2},
		}, {
			name: "",
			nums: []int{3, 2, 1},
		}, {
			name: "",
			nums: []int{2, 3, 1},
		},
		{
			name: "",
			nums: []int{1, 3, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.nums)
			fmt.Println(tt.nums)
		})
	}
}
