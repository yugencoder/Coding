package t13XX

var visited []bool

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}

	adjList := make([][]int, n)
	visited = make([]bool, n)

	for _, conn := range connections {
		adjList[conn[0]] = append(adjList[conn[0]], conn[1])
		adjList[conn[1]] = append(adjList[conn[1]], conn[0])
	}

	noOfComponents := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs2(i, adjList)
			noOfComponents++
		}
	}

	return noOfComponents - 1
}

func dfs2(id int, adjList [][]int) (res int) {
	if visited[id] {
		return
	}

	visited[id] = true
	for _, i := range adjList[id] {
		if !visited[i] {
			dfs2(i, adjList)
		}
	}

	return
}
