package _7XX

import "testing"

func Test_maximumScore(t *testing.T) {
	tests := []struct {
		name string
		a int
		b int
		c int
		want int
	}{
		{
			name: "",
			a:    2,
			b:    4,
			c:    6,
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumScore(tt.a, tt.b, tt.c); got != tt.want {
				t.Errorf("maximumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}