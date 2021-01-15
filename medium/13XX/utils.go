package t13XX

func max(x1 int, x2 int) int {
	if x1 < x2 {
		return x2
	}

	return x1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}
