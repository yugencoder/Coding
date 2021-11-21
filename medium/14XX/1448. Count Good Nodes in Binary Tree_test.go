package _14XX

import (
	"fmt"
	"testing"
)

func Test_goodNodes(t *testing.T) {

	root := new(TreeNode)
	root.Val = 3

	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 5}

	got := goodNodes(root)
	fmt.Println(got)

}
