package X

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func rotateRight(head *ListNode, k int) *ListNode {
	ll := 0

	last := head
	for i := head; i != nil; ll++ {
		last = i
		i = i.Next
	}

	if  ll <= 1 || k%ll == 0{
		return head
	}

	rem := ll - k%ll
	// traverse k
	prev := head
	kth := head
	for i := 0; i < rem; i++ {
		prev = kth
		kth = kth.Next
	}

	// re-arrange
	newHead := kth
	prev.Next = nil
	last.Next = head

	return newHead
}
