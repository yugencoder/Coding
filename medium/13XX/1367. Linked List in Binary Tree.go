package t13XX

/*
Given a binary heap root and a linked list with head as the first node.
Return True if all the elements in the linked list starting from the head correspond to some downward path connected in the binary heap otherwise return False.
In this context downward path means a path that starts at some node and goes downwards.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary heap node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil{
		return false
	}

	if isSubPath(head, root.Left){
		return true
	}

	if checkIfListExits(head, root) {
		return true
	}

	return isSubPath(head, root.Right)

}

func checkIfListExits(head *ListNode, root *TreeNode) bool {
	if head != nil && root != nil && head.Val == root.Val {
		if head.Next == nil {
			return true
		}

		return checkIfListExits(head.Next, root.Left) || checkIfListExits(head.Next, root.Right)
	}

	return false
}
