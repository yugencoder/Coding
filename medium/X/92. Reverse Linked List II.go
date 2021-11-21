package X

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	preHead := new(ListNode)
	preHead.Next = head
	var prev *ListNode
	var lastNode *ListNode

	lCount := 1
	for i := 1; i < left; i++ {
		prev = head
		head = head.Next
		lCount++
	}

	prev.Next, _ = reverseList(head, lCount, left, right, lastNode)

	for node := preHead; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}

	return preHead.Next
}

func reverseList(node *ListNode, i int, left int, right int, lastNode *ListNode) (*ListNode, *ListNode) {
	if node == nil || node.Next == nil || i == right {
		if node != nil && node.Next != nil {
			lastNode = node.Next
		}
		return node, lastNode
	}

	last, llast := reverseList(node.Next, i+1, left, right, lastNode)

	next := node.Next
	next.Next = node
	node.Next = nil

	if i == left {
		node.Next = llast
	}

	return last, llast
}
