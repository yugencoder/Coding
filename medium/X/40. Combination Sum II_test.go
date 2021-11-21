package X

import (
	"fmt"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		//{
		//	name: "",
		//	args: args{
		//		candidates: []int{2, 5, 2, 1, 2},
		//		target:     5,
		//	},
		//	want: nil,
		//},
		{
			name: "",
			args: args{
				candidates: []int{2, 5, 2, 1, 2},
				target:     6,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum2(tt.args.candidates, tt.args.target)
			fmt.Println(got)
		})
	}
}
