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
	var prev, node1, node2 *TreeNode
	var traverse func(root *TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		traverse(root.Left)

		if prev != nil {
			if prev.Val > root.Val {
				if node1 == nil {
					node1 = prev
				}

				node2 = root

			}
		}

		prev = root
		traverse(root.Right)
	}

	traverse(root)

	fmt.Println(node1.Val)
	fmt.Println(node2.Val)
	node1.Val, node2.Val = node2.Val, node1.Val

	return
}

//
//func traverse(root *TreeNode) {
//
//}

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

/*

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
*/
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

//
//func levelOrder(root *TreeNode) {
//	if root == nil {
//		return
//	}
//	var height = heightOfTree(root)
//	fmt.Printf("Height of a tree : %d \n", height)
//	fmt.Printf("Level Order Traversal : \n")
//
//	for i := 1; i <= height; i++ {
//		printLevel(root, i)
//	}
//
//	fmt.Printf("\n")
//}
//*/
