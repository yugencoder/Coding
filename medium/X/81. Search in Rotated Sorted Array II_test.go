package X

import "testing"

func Test_search2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				nums:   []int{2, 5, 6, 0, 0, 1, 2},
				target: 0,
			},
			want: true,
		}, {
			name: "",
			args: args{
				nums:   []int{1, 0, 1, 1, 1},
				target: 0,
			},
			want: true,
		}, {
			name: "",
			args: args{
				nums:   []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1},
				target: 2,
			},
			want: true,
		}, {
			name: "",
			args: args{
				nums:   []int{3, 1, 1},
				target: 3,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				nums:   []int{1, 1, 3},
				target: 3,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				nums:   []int{1, 3, 5},
				target: 1,
			},
			want: true,
		},	{
			name: "",
			args: args{
				nums:   []int{3,1,2,2,2},
				target: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search2(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search2() = %v, want %v", got, tt.want)
			}
		})
	}
}
