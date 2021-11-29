package aiq

import "testing"

func Test_howManyWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				s: "How many eggs are in a half-dozen, 13?",
			},
			want: 7,
		}, {
			name: "",
			args: args{
				s: "he is a good programmer, he won 865 competitions, but sometimes he dont. What do you think? All test-cases should pass. Done-done?",
			},
			want: 21,
		},{
			name: "",
			args: args{
				s: "ds dsaf lkdf kdsa fkldsf, adsbf ldka ads? asd bfdal ds bf[l. akf dhj ds 878 dwa WE DE 7475 dsfh ds RAMU 748 dj.",
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := howManyWords2(tt.args.s); got != tt.want {
				t.Errorf("howManyWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
