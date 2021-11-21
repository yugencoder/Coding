package _8XX

import "testing"

func Test_findTheWinner(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want int
	}{
		{name: "", n: 5, k: 2, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheWinner(tt.n, tt.k); got != tt.want {
				t.Errorf("findTheWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
