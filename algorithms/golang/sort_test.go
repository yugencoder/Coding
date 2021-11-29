package golang

import (
	"testing"
)

func Test_sortAscendingInt(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{
			name: "",
			nums: []int{12, 23, 34, 234, 1, 123, 4, 56, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortAscendingInt(tt.nums)
			sortDescendingInt(tt.nums)
		})
	}
}
