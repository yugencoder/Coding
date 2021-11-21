package X

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	c := 0
	head := &ListNode{}

	for newNode := head; l1 != nil || l2 != nil || c > 0; {
		if l1 != nil {
			newNode.Val = l1.Val
		}
		if l2 != nil {
			newNode.Val += l2.Val
		}

		newNode.Val = (newNode.Val + c) % 10
		c = newNode.Val / 10

		newNode.Next = &ListNode{}
	}

	if c > 0 {

	}

	return head
}
