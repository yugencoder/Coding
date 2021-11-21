package X

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: struct{ s string }{s: "cbbd"}, want: "bb"},
		{name: "", args: struct{ s string }{s: "ccc"}, want: "ccc"},
		{name: "", args: struct{ s string }{s: "caba"}, want: "aba"},
		{name: "", args: struct{ s string }{s: "aacabdkacaa"}, want: "aca"},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
