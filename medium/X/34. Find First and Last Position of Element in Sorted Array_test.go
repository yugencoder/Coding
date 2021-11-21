package X

import (
	"testing"
)

func Test_binaryLeftSearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name            string
		args            args
		want, wantRight int
	}{
		{
			name: "",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			want:      3,
			wantRight: 4,
		},
		{
			name: "",
			args: args{
				nums:   []int{2, 2},
				target: 3,
			},
			want:      -1,
			wantRight: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryLeftSearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("binaryLeftSearch() = %v, want %v", got, tt.want)
			}
			if got := binaryRightSearch(tt.args.nums, tt.args.target); got != tt.wantRight {
				t.Errorf("binaryRightSearch() = %v, want %v", got, tt.wantRight)
			}
		})
	}
}
