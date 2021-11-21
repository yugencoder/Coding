package X

import "testing"

func Test_isInterleave(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				s1: "aabcc",
				s2: "dbbca",
				s3: "aadbbcbcac",
			},
			want: true,
		}, {
			name: "",
			args: args{
				s1: "aabcc",
				s2: "dbbca",
				s3: "aadbbbaccc",
			},
			want: false,
		}, {
			name: "",
			args: args{
				s1: "bcbccabcccbcbbbcbbacaaccccacbaccabaccbabccbabcaabbbccbbbaa",
				s2: "ccbccaaccabacaabccaaccbabcbbaacacaccaacbacbbccccbac",
				s3: "bccbcccabbccaccaccacbacbacbabbcbccbaaccbbaacbcbaacbacbaccaaccabcaccacaacbacbacccbbabcccccbababcaabcbbcccbbbaa",
			},
			want: false,
		}, {
			name: "",
			args: args{
				s1: "aaaaaaaaaaaaaaaaaaaaaaaaaaa",
				s2: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				s3: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			},
			want: false,
		}, {
			name: "",
			args: args{
				s1: "abababababababababababababababababababababababababababababababababababababababababababababababababbb",
				s2: "babababababababababababababababababababababababababababababababababababababababababababababababaaaba",
				s3: "abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababbb",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInterleave(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("isInterleave() = %v, want %v", got, tt.want)
			}
		})
	}
}
