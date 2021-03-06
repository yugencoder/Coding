package _8XX

import "testing"

func Test_maxIceCream(t *testing.T) {
	type args struct {
		costs []int
		coins int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: struct {
			costs []int
			coins int
		}{costs: []int{1,3,2,4,1}, coins: 7}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxIceCream(tt.args.costs, tt.args.coins); got != tt.want {
				t.Errorf("maxIceCream() = %v, want %v", got, tt.want)
			}
		})
	}
}
