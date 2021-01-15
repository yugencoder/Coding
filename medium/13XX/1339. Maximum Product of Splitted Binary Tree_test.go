package t13XX

import "testing"

func Test_maxProduct(t *testing.T) {
	root := new(TreeNode)

	//[1,2,3,4,5,6]
	root.Val = 1
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Right = new(TreeNode)
	root.Right.Val = 3

	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 4
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 5

	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 6

	got := maxProduct(root)
	t.Errorf("maxProduct() = %v", got)

}
