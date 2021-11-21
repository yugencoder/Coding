package _14XX

import "testing"

func Test_maxScore(t *testing.T) {
	type args struct {
		cardPoints []int
		k          int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{name: "", args: args{
		//	cardPoints: []int{1, 2, 3, 4, 5, 6, 1},
		//	k:          3,
		//}, want: 12},
		//{name: "", args: args{
		//	cardPoints: []int{2, 2, 2},
		//	k:          2,
		//},
		//want: 4},
		//{name: "", args: args{
		//	cardPoints: []int{1, 79, 80, 1, 1, 1, 200, 1},
		//	k:          3,
		//}, want: 202},
		//{name: "", args: args{
		//	cardPoints: []int{96,90,41,82,39,74,64,50,30},
		//	k:          8,
		//}, want: 536},
		{name: "", args: args{
			cardPoints: []int{56,27,75,44,38,78,12,43,2,57,71,30,78,38,60,81,61,7,23,85,28,41,47},
			k:          2,
		}, want: 103},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.cardPoints, tt.args.k); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
