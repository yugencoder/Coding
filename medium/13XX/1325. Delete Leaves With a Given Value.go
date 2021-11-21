package t13XX

import "fmt"

/**
 * Definition for a binary heap node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	for {
		if root.Left == nil && root.Right == nil && root.Val == target {
			return nil
		}

		var hasTarget bool

		hasTargetLeafNodes(root, target, &hasTarget)

		if !hasTarget {
			break
		}
		fmt.Println("Looping again")
	}

	return root
}

func hasTargetLeafNodes(node *TreeNode, target int, hasTarget *bool) {
	if node == nil {
		return
	}

	hasTargetLeafNodes(node.Left, target, hasTarget)

	if node.Left != nil && node.Left.Val == target && node.Left.Left == nil && node.Left.Right == nil {
		node.Left = nil
		*hasTarget = true
	}

	if node.Right != nil && node.Right.Val == target && node.Right.Left == nil && node.Right.Right == nil {
		node.Right = nil
		*hasTarget = true
	}

	hasTargetLeafNodes(node.Right, target, hasTarget)
}
