package easy

import "testing"

func Test_validWordAbbreviation(t *testing.T) {
	type args struct {
		word string
		abbr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "", args: args{
			word: "internationalization",
			abbr: "i12iz4n",
		}, want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validWordAbbreviation(tt.args.word, tt.args.abbr); got != tt.want {
				t.Errorf("validWordAbbreviation() = %v, want %v", got, tt.want)
			}
		})
	}
}
