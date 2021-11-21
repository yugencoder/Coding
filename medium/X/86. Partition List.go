package X

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	lHead := new(ListNode)
	rHead := new(ListNode)

	curr := head
	l := lHead
	r := rHead

	for ; curr != nil; curr = curr.Next {
		if curr.Val >= x {
			r.Next = curr
			r = r.Next
		} else {
			l.Next = curr
			l = l.Next
		}
	}

	r.Next = nil
	l.Next = rHead.Next

	return lHead.Next
}
