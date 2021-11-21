package _8XX

import (
	"reflect"
	"testing"
)

func Test_getOrder(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name  string
		tasks [][]int
		want  []int
	}{
		//{name: "", tasks: [][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}}, want: []int{0, 2, 3, 1}},
		//{name: "", tasks: [][]int{{19, 13}, {16, 9}, {21, 10}, {32, 25}, {37, 4}, {49, 24}, {2, 15}, {38, 41}, {37, 34}, {33, 6}, {45, 4}, {18, 18}, {46, 39}, {12, 24}}, want: []int{6, 1, 2, 9, 4, 10, 0, 11, 5, 13, 3, 8, 12, 7}},
		//{name: "", tasks: [][]int{{35, 36}, {11, 7}, {15, 47}, {34, 2}, {47, 19}, {16, 14}, {19, 8}, {7, 34}, {38, 15}, {16, 18}, {27, 22}, {7, 15}, {43, 2}, {10, 5}, {5, 4}, {3, 11}}, want: []int{15, 14, 13, 1, 6, 3, 5, 12, 8, 11, 9, 4, 10, 7, 0, 2}},
		{
			name:  "",
			tasks: [][]int{{35, 36}, {11, 7}, {15, 47}, {34, 2}, {47, 19}, {16, 14}, {19, 8}, {7, 34}, {38, 15}, {16, 18}, {27, 22}, {7, 15}, {43, 2}, {10, 5}, {5, 4}, {3, 11}},
			want:  []int{15, 14, 13, 1, 6, 3, 5, 12, 8, 11, 9, 4, 10, 7, 0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOrder(tt.tasks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
