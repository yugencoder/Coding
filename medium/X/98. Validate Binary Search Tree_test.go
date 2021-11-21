package X

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				root:
				&TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:  4,
						Left: nil,
						Right: &TreeNode{
							Val: 6,
							Left: &TreeNode{
								Val: 3,
							},
							Right: &TreeNode{
								Val: 7,
							},
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
