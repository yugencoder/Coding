package _5XX

import (
	"fmt"
	"testing"
)

func Test_numSubmat(t *testing.T) {
	tests := []struct {
		mat [][]int
	}{
		{
			mat: [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}},
		},
	}

	for _, tt := range tests {
		fmt.Println(numSubmat(tt.mat))
	}
}
