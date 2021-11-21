package t13XX

/**
 * Definition for a binary heap node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type SumNode struct {
	Sum   int
	Left  *SumNode
	Right *SumNode
}

func maxProduct(root *TreeNode) int {
	var res int

	sumRoot := new(SumNode)
	createSumNode(root, sumRoot)

	total := sumRoot.Sum

	getResult(sumRoot.Left, total, &res)
	getResult(sumRoot.Right, total, &res)

	return (res % 1000000007)
}

func getResult(node *SumNode, total int, res *int) {
	if node == nil {
		return
	}

	getResult(node.Left, total, res)
	*res = max(*res, (total - (node.Sum)) * node.Sum)
	getResult(node.Right, total, res)

	return
}

func createSumNode(node *TreeNode, sumNode *SumNode) {
	if node == nil {
		return
	}

	// L
	if node.Left != nil {
		sumNode.Left = new(SumNode)
	}
	createSumNode(node.Left, sumNode.Left)

	if sumNode.Left != nil {
		sumNode.Sum += sumNode.Left.Sum
	}

	// N
	sumNode.Sum += node.Val

	// R
	if node.Right != nil {
		sumNode.Right = new(SumNode)
	}
	createSumNode(node.Right, sumNode.Right)

	if sumNode.Right != nil {
		sumNode.Sum += sumNode.Right.Sum
	}

	return
}
