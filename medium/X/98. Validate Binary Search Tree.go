package X

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return isValidBSTree(root, false, false, math.MinInt32, math.MaxInt32)
}

func isValidBSTree(root *TreeNode, isLeftLimitSet, isRightLimitSet bool, leftLimit int, rightLimit int) bool {
	if root == nil {
		return true
	}

	if root.Left != nil {
		if !(root.Left.Val < root.Val) || isLeftLimitSet && !(root.Left.Val > leftLimit) {
			return false
		}
	}

	if root.Right != nil {
		if !(root.Right.Val > root.Val) || isRightLimitSet && !(root.Right.Val < rightLimit) {
			return false
		}
	}

	return isValidBSTree(root.Left, isLeftLimitSet, true, leftLimit, min(root.Val, rightLimit)) &&
		isValidBSTree(root.Right, true, isRightLimitSet, max(root.Val, leftLimit), rightLimit)
}
