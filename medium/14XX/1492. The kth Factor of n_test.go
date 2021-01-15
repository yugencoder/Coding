package _14XX

import (
	"fmt"
	"testing"
)

func Test_kthFactor(t *testing.T) {

	tests := []struct {
		n int
		k int
	}{
		{n: 1000, k: 3},
	}
	for _, tt := range tests {
		fmt.Println(kthFactor(tt.n, tt.k))
	}
}
