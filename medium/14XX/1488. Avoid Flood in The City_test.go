package _14XX

import (
	"fmt"
	"testing"
)

func Test_avoidFlood(t *testing.T) {
	tests := []struct {
		name  string
		rains []int
	}{
		//{name: "", rains: []int{1, 2, 3, 4}},
		//{name: "", rains: []int{69,0,0,0,69}},
		//{name: "", rains: []int{1,2,0,1,2}},
		//{name: "", rains: []int{0,1,1}},
		{name: "", rains: []int{1,0,2,0,2,1}},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(avoidFlood(tt.rains))
		})
	}
}
