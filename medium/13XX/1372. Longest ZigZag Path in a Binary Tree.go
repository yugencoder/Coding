package t13XX

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
	max := 0
	if root == nil {
		return 0
	}

	findZigZag(root, nil, 0, &max)
	return max
}

func findZigZag(root *TreeNode, isLeft *bool, len int, max *int) {
	if len > *max {
		*max = len
	}

	if root == nil {
		return
	}

	if isLeft == nil {
		isLeft := true
		findZigZag(root.Left, &isLeft, len+1, max)
		isLeft = false
		findZigZag(root.Right, &isLeft, len+1, max)
	} else if *isLeft {
		findZigZag(root.Left, isLeft, 0, max)
		*isLeft = false
		findZigZag(root.Right, isLeft, len+1, max)
	} else {
		findZigZag(root.Right, isLeft, 0, max)
		*isLeft = true
		findZigZag(root.Left, isLeft, len+1, max)
	}
	return
}
