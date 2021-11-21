package _14XX

type TreeNode2 struct {
	val      int
	Next     []int
	hasApple bool
}

func minTime(n int, edges [][]int, hasApple []bool) int {
	treeNodes := make([]TreeNode2, n)
	visited := make([]bool, n)

	for i, apple := range hasApple {
		treeNodes[i] = TreeNode2{
			val:      i,
			Next:     []int{},
			hasApple: apple,
		}
	}

	for _, e := range edges {
		treeNodes[e[0]].Next = append(treeNodes[e[0]].Next, e[1])
		treeNodes[e[1]].Next = append(treeNodes[e[1]].Next, e[0])

	}

	var res int
	for i := 0; i < n; i++ {
		if !visited[i] {
			curRes, _ := dfs(treeNodes, treeNodes[i], 0, visited)
			res += curRes
		}
	}

	return res
}

func dfs(nodes []TreeNode2, node TreeNode2, currTime int, visited []bool) (int, bool) {
	visited[node.val] = true
	subHasApple := false

	if node.hasApple {
		subHasApple = true
	}

	for _, n := range node.Next {
		if !visited[nodes[n].val] {
			endTime, hasApple := dfs(nodes, nodes[n], currTime+1, visited)
			if hasApple {
				subHasApple = hasApple
				currTime = endTime + 1
			}
		}
	}

	return currTime, subHasApple
}
