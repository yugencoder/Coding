package _5XX

import (
	"fmt"
	"testing"
)

func Test_getLastMoment(t *testing.T) {
	tests := []struct {
		n     int
		left  []int
		right []int
	}{
		{n: 9, left: []int{5}, right: []int{4}},
	}
	for _, tt := range tests {
		fmt.Println(getLastMoment(tt.n, tt.left, tt.right))

	}
}
