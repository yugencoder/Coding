package _8XX

import "testing"

func Test_minSideJumps(t *testing.T) {
	tests := []struct {
		name      string
		obstacles []int
		want      int
	}{
		{name: "", obstacles: []int{0, 1, 2, 3, 0}, want: 2},
		{name: "", obstacles: []int{0, 1, 1, 3, 3, 0}, want: 0},
		{name: "", obstacles: []int{0, 2, 1, 0, 3, 0}, want: 2},
		{obstacles: []int{0, 3, 3, 0, 3, 1, 3, 0, 2, 2, 0}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSideJumps(tt.obstacles); got != tt.want {
				t.Errorf("minSideJumps() = %v, want %v", got, tt.want)
			}
		})
	}
}
