package _7XX

import "testing"

func Test_minElements(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		limit int
		goal  int
		want  int
	}{
		{name: "", nums: []int{2, 2, 2, 5, 1, -2}, limit: 5, goal: 126614243, want:25322847},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minElements(tt.nums, tt.limit, tt.goal); got != tt.want {
				t.Errorf("minElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
