package X

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := 0

	for node := head; node != nil; node, l = node.Next, l+1 {
	}

	ns := l - n

	if (ns == 0 || ns == 1) && l == 1 {
		return head.Next
	}

	node := head
	var prev *ListNode
	for i := 0; i < ns && node != nil; i++ {
		prev = node
		node = node.Next
	}

	if prev != nil && node.Next != nil {
		prev.Next = node.Next
	} else {
		prev.Next = nil
	}

	return head
}
