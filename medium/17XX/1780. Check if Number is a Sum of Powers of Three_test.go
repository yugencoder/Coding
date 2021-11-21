package _7XX

import "testing"

func Test_checkPowersOfThree(t *testing.T) {

	tests := []struct {
		name string
		n int
		want bool
	}{
		{name: "", n: 12, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPowersOfThree(tt.n); got != tt.want {
				t.Errorf("checkPowersOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
