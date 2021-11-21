package _7XX

import "testing"

func Test_beautySum(t *testing.T) {
	tests := []struct {
		name string
		s string
		want int
	}{
		{
			name: "",
			s:    "aabcb",
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := beautySum(tt.s); got != tt.want {
				t.Errorf("beautySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
