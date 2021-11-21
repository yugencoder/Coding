package easy

import "container/list"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	stack1 := list.New()
	stack2 := list.New()

	a := headA
	b := headB

	for a != nil || b != nil {
		if a != nil {
			stack1.PushFront(a)
			a = a.Next
		}
		if b != nil {
			stack2.PushFront(b)
			b = b.Next
		}
	}

	var prev *ListNode
	for stack1.Len() != 0 && stack2.Len() != 0 {
		s1 := stack1.Front()
		stack1.Remove(s1)

		s2 := stack2.Front()
		stack2.Remove(s2)

		if s1.Value != s2.Value {
			return prev
		} else {
			prev = s1.Value.(*ListNode)
		}
	}
	return nil
}
