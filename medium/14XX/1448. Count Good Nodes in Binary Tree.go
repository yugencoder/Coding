package _14XX

import (
	"math"
)

/**
 * Definition for a binary heap node.
 * type TreeNode2 struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	count := 0
	max := math.MinInt32
	postorder(root, max, &count)

	return count
}

func postorder(node *TreeNode, max int, count *int) {
	if node == nil {
		return
	}

	prevMax := max

	if node.Val >= max {
		max = node.Val
		*count++
	}

	postorder(node.Left, max, count)
	postorder(node.Right, max, count)

	max = prevMax
}
