package _14XX

import "testing"

func Test_countTriplets(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{arr: []int{2, 3, 1, 6,7}}, want: 4},
		{name: "", args: args{arr: []int{7,11,12,9,5,2,7,17,22}}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTriplets(tt.args.arr); got != tt.want {
				t.Errorf("countTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
