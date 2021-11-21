package _8XX

import "testing"

func Test_reinitializePermutation(t *testing.T) {

	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "", n: 6, want: 4},
		{name: "", n: 4, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reinitializePermutation(tt.n); got != tt.want {
				t.Errorf("reinitializePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
