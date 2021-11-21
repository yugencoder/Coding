package _7XX

import "testing"

func Test_canChoose(t *testing.T) {
	tests := []struct {
		name   string
		groups [][]int
		nums   []int
		want   bool
	}{
		{
			name:   "",
			groups: [][]int{{1, -1, -1}, {3, -2, 0}},
			nums:   []int{1, -1, 0, 1, -1, -1, 3, -2, 0},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canChoose(tt.groups, tt.nums); got != tt.want {
				t.Errorf("canChoose() = %v, want %v", got, tt.want)
			}
		})
	}
}
