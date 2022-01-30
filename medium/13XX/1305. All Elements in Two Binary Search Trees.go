package t13XX

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	r1, r2, res := []int{}, []int{}, []int{}

	getElements(root1, &r1)
	getElements(root2, &r2)

	i, j := 0, 0

	for i < len(r1) && j < len(r2) {
		if r1[i] < r2[j] {
			res = append(res, r1[i])
			i++
		} else {
			res = append(res, r1[j])
			j++
		}
	}

	for i < len(r1) {
		res = append(res, r1[i])
		i++
	}

	for i < len(r2) {
		res = append(res, r2[i])
		i++
	}

}

func getElements(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		getElements(root.Left, res)
	}

	*res = append(*res, root.Val)

	if root.Right != nil {
		getElements(root.Right, res)
	}
}
