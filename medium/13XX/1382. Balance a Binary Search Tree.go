package t13XX

import "math"

/*
Given a binary search heap, return a balanced binary search heap with the same node values.

A binary search heap is balanced if and only if the depth of the two subtrees of every node never differ by more than 1.

If there is more than one answer, return any of them.



Example 1:



Input: root = [1,null,2,null,3,null,4,null,null]
Output: [2,1,3,null,null,null,4]
Explanation: This is not the only correct answer, [3,1,4,null,2,null,null] is also correct.

*/

func BalanceBST(root *TreeNode) *TreeNode {
	sortedList := []int{}

	traverseInorder(root, &sortedList)
	return buildBalancedTree(0, len(sortedList)-1, sortedList)
}

func buildBalancedTree(left, right int, list []int) *TreeNode {
	if left > right {
		return nil
	}
	if left == right {
		return &TreeNode{list[left], nil, nil}
	}

	middle := left + int(math.Floor(float64(right-left))/2)
	root := new(TreeNode)
	root.Val = list[middle]

	root.Left = buildBalancedTree(left, middle-1, list)
	root.Right = buildBalancedTree(middle+1, right, list)

	return root
}

// todo - if append or structure changed in any way - slice is not a reference :)
func traverseInorder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		traverseInorder(root.Left, list)
	}
	*list = append(*list, root.Val)

	if root.Right != nil {
		traverseInorder(root.Right, list)
	}
}
