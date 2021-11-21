package _14XX

import "testing"

func Test_checkIfCanBreak(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		//{name: "", args: args{
		//	s1: "leetcodee",
		//	s2: "interview",
		//}, want: true},
		//{name: "", args: args{
		//	s1: "abe",
		//	s2: "acd",
		//}, want: false},
		{name: "", args: args{
			s1: "abc",
			s2: "xya",
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfCanBreak(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkIfCanBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
