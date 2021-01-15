package t13XX

import (
	"fmt"
	"testing"
)

func Test_xorQueries(t *testing.T) {
	tests := []struct {
		arr     []int
		queries [][]int
	}{
		{
			arr:     []int{1, 3, 4, 8},
			queries: [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}},
	}
	for _, tt := range tests {
		fmt.Println(xorQueries(tt.arr, tt.queries))
	}
}
