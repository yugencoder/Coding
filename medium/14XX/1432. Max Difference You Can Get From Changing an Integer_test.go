package _14XX

import "testing"

func Test_maxDiff(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{num: 1101057}, want: 8808050},
		//{name: "", args: args{num: 9}, want: 8},
		//{name: "", args: args{num: 555}, want: 888},
		//{name: "", args: args{num: 123456}, want: 820000},
		//{name: "", args: args{num: 9288}, want: 8700},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDiff(tt.args.num); got != tt.want {
				t.Errorf("maxDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
