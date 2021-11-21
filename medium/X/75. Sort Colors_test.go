package X

import (
	"fmt"
	"testing"
)

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				nums: []int{2, 0, 2, 1, 1, 0},
			},
		},
		{
			name: "",
			args: args{
				nums: []int{1,0,0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := tt.args.nums
			sortColors(nums)
			fmt.Println(nums)
		})
	}
}
