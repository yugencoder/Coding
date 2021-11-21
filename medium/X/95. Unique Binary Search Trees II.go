package X

import "fmt"

//
///**
// * Definition for a binary tree node.
// * type TreeNode struct {
// *     Val int
// *     Left *TreeNode
// *     Right *TreeNode
// * }
// */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	all := make([]int, n)
	for i := 1; i <= n; i++ {
		all[i-1] = i
	}

	r := generateTree(all)
	fmt.Println(r)

	return r
}

func generateTree(val []int) []*TreeNode {
	res := []*TreeNode{}

	if len(val) == 0 {
		return nil
	}
	if len(val) == 1 {
		return []*TreeNode{
			{
				Val: val[0],
			},
		}
	}

	for i, v := range val {
		left := generateTree(val[:i])
		right := generateTree(val[i+1:])

		if len(left) == 0 {
			left = []*TreeNode{nil}
		} else if len(right) == 0 {
			right = []*TreeNode{nil}
		}

		for _, l := range left {
			for _, r := range right {
				node := new(TreeNode)
				node.Val = v
				node.Left = l
				node.Right = r
				res = append(
					res,
					node,
				)
			}
		}

	}

	return res
}
