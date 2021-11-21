package _7XX

import (
	"reflect"
	"testing"
)

func Test_highestPeak(t *testing.T) {

	tests := []struct {
		name    string
		isWater [][]int
		want    [][]int
	}{
		{name: "", isWater: [][]int{{0, 1}, {0, 0}}, want: [][]int{{1, 0}, {2, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := highestPeak(tt.isWater); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("highestPeak() = %v, want %v", got, tt.want)
			}
		})
	}
}
