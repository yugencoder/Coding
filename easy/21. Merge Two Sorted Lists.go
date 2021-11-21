package easy

//
//* Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var prev *ListNode
	root := l1

	if l1 == nil || l2 == nil {
		if l1 == nil {
			return l2
		} else {
			return l1
		}
	}

	for l2 != nil && l1 != nil {
		if l1.Val <= l2.Val {
			prev = l1
			l1 = l1.Next
		} else {
			if prev != nil {
				prev.Next = l2
				nextL2 := l2.Next
				l2.Next = l1

				l2 = nextL2
			} else {
				next := l2.Next
				l2.Next = l1
				root = l2

				prev = l2
				l2 = next
			}
		}
	}

	if l2 != nil {
		prev.Next = l2
	}

	return root
}
