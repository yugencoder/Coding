package aiq

import (
	"reflect"
	"testing"
)

func Test_distanceMetric(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 2, 1, 1, 2, 3},
			},
			want: []int{5, 3, 3, 4, 3, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distanceMetric2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distanceMetric() = %v, want %v", got, tt.want)
			}
		})
	}
}