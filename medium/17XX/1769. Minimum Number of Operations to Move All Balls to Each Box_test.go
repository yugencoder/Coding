package _7XX

import (
	"reflect"
	"testing"
)

func Test_minNoOperations(t *testing.T) {

	tests := []struct {
		name  string
		boxes string
		want  []int
	}{
		{name: "", boxes: "110", want: []int{1, 1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNoOperations(tt.boxes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minNoOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
