package X

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
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

func recoverTree(root *TreeNode) {

	var firstNode, secondNode, prevNode *TreeNode
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if prevNode != nil {
			if node.Val < prevNode.Val {
				if firstNode == nil {
					firstNode = prevNode
				}
				secondNode = node
			}
		}
		prevNode = node
		inorder(node.Right)
	}
	inorder(root)
	firstNode.Val, secondNode.Val = secondNode.Val, firstNode.Val
}

func heightOfTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var Left = heightOfTree(root.Left)
	var Right = heightOfTree(root.Right)

	if Left > Right {
		return Left + 1
	} else {
		return Right + 1
	}
}

func printLevel(root *TreeNode, h int) {
	if root == nil {
		return
	}

	if h == 1 {
		fmt.Printf("%d ", root.Val)
		return
	}
	if h > 1 {
		printLevel(root.Left, h-1)

		printLevel(root.Right, h-1)
	}
}

func levelOrder(root *TreeNode) {
	if root == nil {
		return
	}
	var height = heightOfTree(root)
	fmt.Printf("Height of a tree : %d \n", height)
	fmt.Printf("Level Order Traversal : \n")

	for i := 1; i <= height; i++ {
		printLevel(root, i)
	}

	fmt.Printf("\n")
}
