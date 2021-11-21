package _14XX

import "testing"

func Test_longestSubarray(t *testing.T) {
	type args struct {
		nums  []int
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{
			nums:  []int{1, 5, 6, 7, 8, 10, 6, 5, 6},
			limit: 4,
		}, want: 5},
		//{name: "", args: args{
		//	nums:  []int{10, 1, 2, 4, 7, 2},
		//	limit: 5,
		//}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.args.nums, tt.args.limit); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
