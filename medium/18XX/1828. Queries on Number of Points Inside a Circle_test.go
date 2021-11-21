package _8XX

import (
	"reflect"
	"testing"
)

func Test_countPoints(t *testing.T) {
	type args struct {

	}
	tests := []struct {
		name string
		points  [][]int
		queries [][]int
		want []int
	}{
		{name: "", points:[][]int{{1,1},{2,2},{3,3},{4,4},{5,5}} , queries:[][]int{{1,2,2},{2,2,2},{4,3,2},{4,3,3}} , want: []int{2,3,2,4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPoints(tt.points, tt.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
