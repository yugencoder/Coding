package X

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	preHead := &ListNode{}

	prev := preHead
	prev.Next = head

	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next
		second := first.Next
		next := second.Next

		second.Next = first
		first.Next = next
		prev.Next = second

		prev = first
	}

	return preHead.Next
}
