package X

import (
	"reflect"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: struct{ head *ListNode }{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					2,
					&ListNode{
						Val: 3,
						Next: &ListNode{4, nil,
						}},
				},
			},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					1,
					&ListNode{
						Val: 4,
						Next: &ListNode{3, nil,
						}},
				},
			},

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
