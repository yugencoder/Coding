package IQ

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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := howManyWords(tt.args.s); got != tt.want {
				t.Errorf("howManyWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
