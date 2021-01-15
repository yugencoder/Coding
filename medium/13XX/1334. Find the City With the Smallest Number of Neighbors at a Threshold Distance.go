package t13XX

import "math"

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	dis := make([][]int, n)

	// easier way ?
	for i := 0; i < n; i++ {

	}

	for i := 0; i < n; i++ {
		dis[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				dis[i][i] = 0
				continue
			}
			dis[i][j] = math.MaxInt32
		}
	}

	// create list
	//adjList := make([][]int, 100, 100)
	for _, e := range edges {
		//adjList[e[0]][e[1]] = e[2]
		dis[e[0]][e[1]] = e[2]
		dis[e[1]][e[0]] = e[2]
	}
	// apply floyd marhsals
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dis[i][j] = min(dis[i][j], dis[i][k]+dis[k][j])
			}
		}
	}

	// find above threshold values
	minMap := map[int]int{}
	minVal := math.MaxInt32
	minValIndex := -1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dis[i][j] <= distanceThreshold{
				minMap[i] += 1
			}
		}
		if minMap[i] > 0 {
			minVal = min(minVal, minMap[i])
			if minVal == minMap[i] {
				minValIndex = i
			}
		}
	}
	return minValIndex
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}
