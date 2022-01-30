package _XX

import "testing"

func Test_findMaximumXOR(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				[]int{3, 10, 5, 25, 2, 8},
			},
			want: 28,
		},
		{
			name: "",
			args: args{
				[]int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70},
			},
			want: 127,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaximumXOR(tt.args.nums); got != tt.want {
				t.Errorf("findMaximumXOR() = %v, want %v", got, tt.want)
			}
		})
	}
}
