package X

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: struct{ s string }{s: "   -42"}, want: -42},
		{name: "", args: struct{ s string }{s: "4193 with words"}, want: 4193},
		{name: "", args: struct{ s string }{s: "-91283472332"}, want: -2147483648},
		{name: "", args: struct{ s string }{s: "words and 987"}, want: 0},
		{name: "", args: struct{ s string }{s: "3.14159"}, want: 3},
		{name: "", args: struct{ s string }{s: "+-12"}, want: 0},
		{name: "", args: struct{ s string }{s: "00000-42a1234"}, want: 0},
		{name: "", args: struct{ s string }{s: "   +0 123"}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
