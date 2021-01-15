package t13XX

import (
	"fmt"
	"testing"
)

func Test_diagonalSort(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		want [][]int
	}{
		{
			name: "",
			mat:  [][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diagonalSort(tt.mat)
			fmt.Println(got)
		})
	}
}
