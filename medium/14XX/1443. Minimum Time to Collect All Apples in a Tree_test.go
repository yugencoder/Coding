package _14XX

import "testing"

func Test_minTime(t *testing.T) {
	type args struct {
		n        int
		edges    [][]int
		hasApple []bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{
			n:        7,
			edges:    [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
			hasApple: []bool{false, false, true, false, true, true, false},
		}, want: 8},
		{name: "", args: args{
			n:        7,
			edges:    [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
			hasApple: []bool{false, false, true, false, false, true, false},
		}, want: 6},
		{name: "", args: args{
			n:        4,
			edges:    [][]int{{0, 2}, {0, 3}, {1, 2}},
			hasApple: []bool{false, true, false, false},
		}, want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTime(tt.args.n, tt.args.edges, tt.args.hasApple); got != tt.want {
				t.Errorf("minTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
