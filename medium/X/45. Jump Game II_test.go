package X

import "testing"

func Test_jump(t *testing.T) {
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
				nums: []int{2,3,1,1,4},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums: []int{1},
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				nums: []int{1,2},
			},
			want: 1,
		},		{
			name: "",
			args: args{
				nums: []int{7,0,9,6,9,6,1,7,9,0,1,2,9,0,3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
