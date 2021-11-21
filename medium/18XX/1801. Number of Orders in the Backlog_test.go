package _8XX

import "testing"

func Test_getNumberOfBacklogOrders(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name   string
		orders [][]int
		want   int
	}{
		{
			name:   "",
			orders: [][]int{{10, 5, 0}, {15, 2, 1}, {25, 1, 1}, {30, 4, 0}},
			want:   6,
		},		{
			name:   "",
			orders: [][]int{{7,1000000000,1},{15,3,0},{5,999999995,0},{5,1,1}},
			want:   999999984,
		},	{
			name:   "",
			orders: [][]int{{26,7,0},{16,1,1},{14,20,0},{23,15,1},{24,26,0},{19,4,1},{1,1,0}},
			want:   34,
		},{
			name:   "",
			orders: [][]int{{1,29,1},{22,7,1},{24,1,0},{25,15,1},{18,8,1},{8,22,0},{25,15,1},{30,1,1},{27,30,0}},
			want:   34,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumberOfBacklogOrders(tt.orders); got != tt.want {
				t.Errorf("getNumberOfBacklogOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}
