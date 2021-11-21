package _14XX

import (
	"reflect"
	"testing"
)

func Test_simplifiedFractions(t *testing.T) {

	tests := []struct {
		name string
		args int
		want []string
	}{
		{name: "", args: 4, want: []string{"1/2", "1/3", "1/4", "2/3", "3/4"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifiedFractions(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("simplifiedFractions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gcd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "", args: args{
			a: 6,
			b: 3,
		}, want:3 },
		{name: "", args: args{
			a: 4,
			b: 35,
		}, want:1 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}