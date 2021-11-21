package _7XX

import "testing"

func Test_countRestrictedPaths(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  int
	}{
		{
			name:  "",
			n:     5,
			edges: [][]int{{1, 2, 3}, {1, 3, 3}, {2, 3, 1}, {1, 4, 2}, {5, 2, 2}, {3, 5, 1}, {5, 4, 10}},
			want:  3,
		},
		{
			name:  "",
			n:     5,
			edges: [][]int{{2, 1, 80634}, {3, 4, 76993}, {1, 5, 49551}, {4, 2, 10047}, {2, 3, 74814}, {4, 1, 72126}},
			want:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRestrictedPaths(tt.n, tt.edges); got != tt.want {
				t.Errorf("countRestrictedPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
