package _14XX

import "testing"

func Test_getHappyString(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{
			n: 1,
			k: 3,
		}, want:"c" },
		{name: "", args: args{
			n: 3,
			k: 9,
		}, want:"cab" },
		{name: "", args: args{
			n: 1,
			k: 4,
		}, want:"" },

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHappyString(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("getHappyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
