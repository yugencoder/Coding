package t13XX

import (
	"fmt"
	"testing"
)

func Test_matrixBlockSum(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		K    int
		want [][]int
	}{
		// TODO: Add test cases.
		//[[1,2,3],[4,5,6],[7,8,9]
		{
			mat: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			K:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(matrixBlockSum(tt.mat, tt.K))
		})
	}
}
