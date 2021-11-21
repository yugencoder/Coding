package _8XX

import "testing"

func Test_maxValue(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		index  int
		maxSum int
		want   int
	}{
		{
			name:   "",
			n:      4,
			index:  2,
			maxSum: 6,
			want:   2,
		},
		{
			name:   "",
			n:      6,
			index:  1,
			maxSum: 10,
			want:   3,
		},
		{
			name:   "",
			n:      3,
			index:  2,
			maxSum: 18,
			want:   7,
		},
		{
			name:   "",
			n:      8,
			index:  7,
			maxSum: 14,
			want:   4,
		},
		{
			name:   "",
			n:      4,
			index:  0,
			maxSum: 4,
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValue(tt.n, tt.index, tt.maxSum); got != tt.want {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
