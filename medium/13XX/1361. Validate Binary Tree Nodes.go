package t13XX

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {

	roots := findRoots(n, leftChild, rightChild)
	if len(roots) > 1 || len(roots) == 0 {
		return false
	}

	visitedMap := map[int]bool{}
	valid := true
	traverseTree(roots[0], n, leftChild, rightChild, &valid, &visitedMap, 0)
	if valid && len(visitedMap) == n {
		return true
	}

	return false
}

func findRoots(n int, leftChild []int, rightChild []int) []int {
	nonRootNodes := map[int]bool{}
	roots := []int{}

	for i, l := range leftChild {
		if l != -1 {
			nonRootNodes[l] = true
		}

		if rightChild[i] != -1 {
			nonRootNodes[rightChild[i]] = true
		}
	}

	for i := 0; i < n; i++ {
		if !nonRootNodes[i] {
			roots = append(roots, i)
		}
	}

	return roots
}

func traverseTree(i, n int, leftChild []int, rightChild []int, valid *bool, visitedMap *map[int]bool, count int) {
	count++
	if !*valid || (*visitedMap)[i] || count > n {
		*valid = false
		return
	}

	if leftChild[i] != -1 {
		traverseTree(leftChild[i], n, leftChild, rightChild, valid, visitedMap, count)
	}

	(*visitedMap)[i] = true

	if rightChild[i] != -1 {
		traverseTree(rightChild[i], n, leftChild, rightChild, valid, visitedMap, count)
	}
}
