package _14XX

import "testing"

func Test_findMinFibonacciNumbers(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{k: 10}, want: 2},
		{name: "", args: args{k: 19}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinFibonacciNumbers(tt.args.k); got != tt.want {
				t.Errorf("findMinFibonacciNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}