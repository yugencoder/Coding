package _7XX

func findCenter(edges [][]int) int {
	if len(edges) == 0 {
		return edges[0][0]
	}

	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	}

	return edges[0][1]
}
