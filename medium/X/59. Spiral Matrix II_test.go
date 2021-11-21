package X

import (
	"reflect"
	"testing"
)

func Test_generateMatrix(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]int
	}{
		{
			name: "",
			n:    3,
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMatrix(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
