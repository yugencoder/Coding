package _14XX

import "testing"

func Test_findLeastNumOfUniqueInts(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		k    int
		want int
	}{
		{name: "", arr: []int{5, 5, 4}, k: 1, want: 1},
		//{name: "", arr: []int{5, 5, 4}, k: 3, want: 1},
		//{name: "", arr: []int{4, 3, 1, 1, 3, 3, 2}, k: 3, want: 2},
		//{name: "", arr: []int{4, 3, 1, 1, 3, 3, 2}, k: 3, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLeastNumOfUniqueInts(tt.arr, tt.k); got != tt.want {
				t.Errorf("findLeastNumOfUniqueInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
