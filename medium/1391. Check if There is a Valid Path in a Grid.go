package problem

import "strings"

func find(root []int, a int) int {
	if root[a] == a {
		return a
	}
	root[a] = find(root, root[a])
	return root[a]
}

func union(root []int, a, b int) {
	rootB := find(root, b)
	rootA := find(root, a)
	if rootA != rootB {
		root[rootB] = rootA
	}

}

func hasValidPath2(grid [][]int) bool {
	dirMap := map[int]string{
		1: "lr",
		2: "ud",
		3: "ld",
		4: "rd",
		5: "ul",
		6: "ur",
	}

	coordMap := map[string][]int{
		"l": []int{0, -1},
		"r": []int{0, 1},
		"u": []int{-1, 0},
		"d": []int{1, 0},
	}

	next := map[string]string{
		"l": "r",
		"r": "l",
		"u": "d",
		"d": "u",
	}

	i, j := 0, 0
	expected := ""

	switch grid[i][j] {
	case 1, 3:
		expected = "l"
	case 2, 6:
		expected = "u"
	case 4:
		if hasPath(i, grid, j, dirMap, "r", coordMap, next) {
			return true
		} else {
			return hasPath(i, grid, j, dirMap, "d", coordMap, next)
		}
	default:
		return false
	}

	return hasPath(i, grid, j, dirMap, expected, coordMap, next)
}

func hasPath(i int, grid [][]int, j int, dirMap map[int]string, expected string, coordMap map[string][]int, next map[string]string) bool {
	for i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) {
		currDir := dirMap[grid[i][j]]

		if strings.Contains(currDir, expected) {
			if i == len(grid)-1 && j == len(grid[0])-1 {
				return true
			}

			currDir = strings.Replace(currDir, expected, "", 1)
			i = i + coordMap[currDir][0]
			j = j + coordMap[currDir][1]

			expected = next[currDir]
		} else {
			return false
		}
	}

	return false
}
