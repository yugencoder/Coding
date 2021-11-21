package _8XX

import "testing"

func Test_getGreaterThanEqualTo(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		nums  [][]int
		start int
		end   int
		want  int
	}{
		{
			name:  "",
			n:     5,
			nums:  [][]int{{0}, {1}, {4}, {5}, {6}, {7}, {23}, {78}, {123}},
			start: 0,
			end:   8,
			want:  3,
		},
		{
			name:  "",
			n:     8,
			nums:  [][]int{{1}, {4}, {5}, {6}, {7}, {23}, {78}, {123}},
			start: 0,
			end:   8,
			want:  5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGreaterThanEqualTo(tt.n, tt.nums, tt.start, tt.end); got != tt.want {
				t.Errorf("getGreaterThanEqualTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minAbsoluteSumDiff(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  int
	}{
		{
			name:  "",
			nums1: []int{10,20,30,40,50},
			nums2: []int{510,520,530,540,550},
			want:  3,
		},{
			name:  "",
			nums1: []int{1, 7, 5},
			nums2: []int{2, 3, 5},
			want:  3,
		},
		{
			name:  "",
			nums1: []int{2, 4, 6, 8, 10},
			nums2: []int{2, 4, 6, 8, 10},
			want:  0,
		},
		{
			name:  "",
			nums1: []int{1, 10, 4, 4, 2, 7},
			nums2: []int{9, 3, 5, 1, 7, 4},
			want:  20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAbsoluteSumDiff(tt.nums1, tt.nums2); got != tt.want {
				t.Errorf("minAbsoluteSumDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
