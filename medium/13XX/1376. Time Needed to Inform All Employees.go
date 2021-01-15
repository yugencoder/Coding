package t13XX

import "math"

func NumOfMinutes(n int, headID int, manager []int, informTime []int) int {
	adjList := make([][]int, n)
	for i := 0; i < n; i++ {
		if manager[i] == -1 {
			continue
		}

		if adjList[manager[i]] == nil {
			adjList[manager[i]] = []int{}
		}

		adjList[manager[i]] = append(adjList[manager[i]], i)

	}

	return dfs(headID, adjList, informTime)
}

func dfs(id int, adjList [][]int, informTime []int) (res int) {
	if len(adjList[id]) == 0 {
		return
	}

	res = math.MinInt32

	for _, val := range adjList[id] {
		res = max(res, informTime[id]+dfs(val, adjList, informTime))
	}

	return
}
