package X

import "testing"

func Test_recoverTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				&TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   4,
						Right: nil,
						Left: &TreeNode{
							Val:   2,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
		}, {
			name: "",
			args: args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  3,
						Left: nil,
						Right: &TreeNode{
							Val:   2,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//levelOrder(tt.args.root)
			recoverTree(tt.args.root)
			//levelOrder(tt.args.root)
		})
	}
}
