package t13XX

import (
	"fmt"
	"testing"
)

func Test_makeConnected(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		connections [][]int
	}{
		//{name: "", n: 4, connections: [][]int{{0, 1}, {0, 2}, {1, 2}}},
		{name: "", n: 6, connections: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeConnected(tt.n, tt.connections)
			fmt.Println(got)

		})
	}
}
