package _5XX

import (
	"fmt"
	"testing"
)

func Test_minDifference(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "", args: args{
			nums: []int{1, 5, 0, 10, 14},
		},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minDifference(tt.args.nums)
			fmt.Println(got)
		})
	}
}
