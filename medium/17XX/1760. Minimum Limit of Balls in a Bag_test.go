package _7XX

import "testing"

func Test_minimumSize(t *testing.T) {
	tests := []struct {
		name          string
		nums          []int
		maxOperations int
		want          int
	}{
		{name: "", nums: []int{9}, maxOperations: 2, want: 3},

		{name: "", nums: []int{2, 4, 8, 2}, maxOperations: 4, want: 2},
		{name: "", nums: []int{7, 17}, maxOperations: 2, want: 7},
		{name: "", nums: []int{9}, maxOperations: 2, want: 3},
		{name: "", nums: []int{501, 67, 484, 937, 816, 895, 294, 240, 736, 245, 266, 698, 371, 538, 265, 309, 422, 476, 827, 816, 927, 379, 732, 941, 119, 303, 914, 311, 518, 843, 359, 198, 341, 633, 671, 22, 23, 235, 556, 92, 239, 389, 393, 74, 799, 792, 477, 696, 150, 39, 979, 97, 330, 81, 798, 630, 954, 955, 633, 438, 342, 909, 103, 82, 184, 240, 693, 705, 225, 55, 311, 181, 422, 371, 177, 156, 138, 806, 255, 633, 801, 427, 214, 49, 957, 738, 325, 317, 576, 117, 955, 931, 83, 100, 931, 227, 616, 109, 333, 285, 113, 744, 398, 981, 760}, maxOperations: 92, want: 333},


		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSize(tt.nums, tt.maxOperations); got != tt.want {
				t.Errorf("minimumSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
