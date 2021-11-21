package X

import (
	"fmt"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}

	got := addTwoNumbers(l1, l2)
	fmt.Println(got)
}
