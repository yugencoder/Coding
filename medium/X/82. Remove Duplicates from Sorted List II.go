package X

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	preHead := new(ListNode)
	preHead.Next = head

	prev := preHead
	curr := head

	for curr != nil {
		next := curr.Next
		if next == nil {
			return preHead.Next
		}

		if curr.Val == next.Val {
			for ; next != nil && curr.Val == next.Val; next = curr.Next {
				curr = curr.Next
			}
			prev.Next = curr.Next
		} else {
			prev = prev.Next
		}

		curr = curr.Next
	}

	return preHead.Next
}
