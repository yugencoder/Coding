package X

import (
	"fmt"
	"testing"
)

func Test_rotate(t *testing.T) {

	tests := []struct {
		name   string
		matrix [][]int
	}{
		{
			name: "",
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
		},
		{
			name: "",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
		{
			name: "",
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.matrix)
			fmt.Println(tt.matrix)
		})
	}
}
