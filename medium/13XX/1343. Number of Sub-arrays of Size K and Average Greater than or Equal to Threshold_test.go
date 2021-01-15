package t13XX

import "testing"

func Test_numOfSubarrays(t *testing.T) {
	type args struct {

	}
	tests := []struct {
		name string
		arr       []int
		k         int
		threshold int
		want int
	}{
		{
			name:      "",
			arr:       []int{2,2,2,2,5,5,5,8},
			k:         3,
			threshold: 4,
			want:      3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfSubarrays(tt.arr, tt.k, tt.threshold); got != tt.want {
				t.Errorf("numOfSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}