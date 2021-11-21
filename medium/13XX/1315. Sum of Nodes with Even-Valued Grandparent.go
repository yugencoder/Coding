package t13XX

/**
 * Definition for a binary heap node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumEvenGrandparent(root *TreeNode) int {
	return traverse(root, -1, -1)
}

func traverse(node *TreeNode, parent int, gp int) int {
	if node == nil {
		return 0
	}

	sum := traverse(node.Left, node.Val, parent)

	if gp >= 0 && gp%2 == 0 {
		sum += node.Val
	}

	sum += traverse(node.Right, node.Val, parent)

	return sum
}
