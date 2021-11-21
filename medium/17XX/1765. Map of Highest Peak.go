package _7XX

var coordMap map[int][][]int

func highestPeak(isWater [][]int) [][]int {
	m := len(isWater)
	n := len(isWater[0])
	coordMap = map[int][][]int{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				isWater[i][j] = 0
				coordMap[0] = append(coordMap[0], []int{i, j})
			} else {
				isWater[i][j] = -1
			}
		}
	}

	for i := 0; ; i++ {
		coords, ok := coordMap[i]
		if !ok {
			return isWater
		}
		for _, coord := range coords {
			fillAdjacent(isWater, coord[0], coord[1], i+1)
		}
	}

	return isWater
}

func fillAdjacent(isWater [][]int, i int, j int, val int) {
	incs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for _, inc := range incs {
		i_next := i + inc[0]
		j_next := j + inc[1]

		if (i_next >= 0 && i_next < len(isWater)) && (j_next >= 0 && j_next < len(isWater[0])) {
			if isWater[i_next][j_next] == -1 {
				isWater[i_next][j_next] = val
				coordMap[val] = append(coordMap[val], []int{i_next, j_next})
			}
		}
	}
}

