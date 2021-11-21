package X

import "testing"

func Test_search(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},

		{
			name:   "",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},

		{
			name:   "",
			nums:   []int{1, 3},
			target: 2,
			want:   -1,
		},
		{
			name:   "",
			nums:   []int{1, 3},
			target: 3,
			want:   1,
		},
		{
			name:   "",
			nums:   []int{5, 1, 3},
			target: 2,
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.nums, tt.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
