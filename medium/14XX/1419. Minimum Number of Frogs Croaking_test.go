package _14XX

import "testing"

func Test_minNumberOfFrogs(t *testing.T) {
	type args struct {
		croakOfFrogs string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{croakOfFrogs: "croakcroak"}, want: 1},
		{name: "", args: args{croakOfFrogs: "crcoakroak"}, want: 2},
		{name: "", args: args{croakOfFrogs: "ccrrooaakk"}, want: 2},
		{name: "", args: args{croakOfFrogs: "crocakcroraoakk"}, want: 2},


	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumberOfFrogs(tt.args.croakOfFrogs); got != tt.want {
				t.Errorf("minNumberOfFrogs() = %v, want %v", got, tt.want)
			}
		})
	}
}
