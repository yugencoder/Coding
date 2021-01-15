package _14XX

import (
	"fmt"
	"testing"
)

func Test_numSubseq(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
	}{
		//{nums: []int{3, 5, 6, 7}, target: 9},
		//{nums: []int{3, 3, 6, 8}, target: 10},
		//{nums: []int{2, 3, 3, 4, 6, 7}, target: 12},
		//{nums: []int{5, 2, 4, 1, 7, 6, 8}, target: 16},
		//{nums: []int{7, 10, 7, 3, 7, 5, 4}, target: 12},
		{nums: []int{27, 21, 14, 2, 15, 1, 19, 8, 12, 24, 21, 8, 12, 10, 11, 30, 15, 18, 28, 14, 26, 9, 2, 24, 23, 11, 7, 12, 9, 17, 30, 9, 28, 2, 14, 22, 19, 19, 27, 6, 15, 12, 29, 2, 30, 11, 20, 30, 21, 20, 2, 22, 6, 14, 13, 19, 21, 10, 18, 30, 2, 20, 28, 22},
			target: 31},
	}
	for _, tt := range tests {
		fmt.Println(numSubseq(tt.nums, tt.target))
	}
}
