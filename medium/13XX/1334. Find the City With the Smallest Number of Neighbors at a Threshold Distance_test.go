package t13XX

import (
	"fmt"
	"testing"
)

func Test_findTheCity(t *testing.T) {
	tests := []struct {
		n                 int
		edges             [][]int
		distanceThreshold int
	}{
		//{
		//	4,
		//	[][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}},
		//	4,
		//},
		{
			5,
			[][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}},
			2,
		},
		//5
		//[[0,1,2],[0,4,8],[1,2,3],[1,4,2],[2,3,1],[3,4,1]]
		//2
	}
	for _, tt := range tests {
		got := findTheCity(tt.n, tt.edges, tt.distanceThreshold)
		fmt.Println(got)

	}
}
