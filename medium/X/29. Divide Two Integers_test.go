package X

import "testing"

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{10, 3}, want: 3},
		{name: "", args: args{7, -3}, want: -2},
		{name: "", args: args{-1, 1}, want: -1},
		{name: "", args: args{-1, -1}, want: 1},

		{name: "", args: args{-2147483648, -1}, want: 2147483647},
		{name: "", args: args{2147483647, 1}, want: 2147483647},
		{name: "", args: args{-2147483648, 1}, want: -2147483648},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
