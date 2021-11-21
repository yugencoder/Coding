package _7XX

import "testing"

func Test_closestCost(t *testing.T) {
	tests := []struct {
		name string
		baseCosts    []int
		toppingCosts []int
		target       int
		want int
	}{
		{
			name:         "",
			baseCosts:    []int{2, 3},
			toppingCosts: []int{4,5,100},
			target:       18	,
			want:         17,
		},		{
			name:         "",
			baseCosts:    []int{4},
			toppingCosts: []int{9},
			target:       9	,
			want:         13,
		},	{
			name:         "",
			baseCosts:    []int{9,10,1},
			toppingCosts: []int{1,8,8,1,1,8},
			target:       8	,
			want:         7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestCost(tt.baseCosts, tt.toppingCosts, tt.target); got != tt.want {
				t.Errorf("closestCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
