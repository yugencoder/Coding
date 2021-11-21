package X

import "math"

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(newInterval) == 0 || len(intervals) == 0 {
		return append(intervals, [][]int{newInterval}...)
	}

	if intervals[len(intervals)-1][1] < newInterval[0] {
		return append(intervals, newInterval)
	} else if newInterval[1] < intervals[0][0] {
		return append([][]int{newInterval}, intervals...)
	}

	res := [][]int{}
	nx := newInterval[0]
	ny := newInterval[1]

	start := math.MaxInt32
	end := -1

	for i := 0; i < len(intervals); i++ {
		x := intervals[i][0]
		y := intervals[i][1]

		// intersection case
		if (nx >= x && nx <= y || x >= nx && x <= ny) || (ny >= x && ny <= y || y >= nx && y <= ny) {
			for ; i < len(intervals) && ((nx >= x && nx <= y || x >= nx && x <= ny) || (ny >= x && ny <= y || y >= nx && y <= ny)); i++ {
				start = min(start, min(nx, x))
				end = max(end, max(ny, y))

				if i < len(intervals)-1 {
					x = intervals[i+1][0]
					y = intervals[i+1][1]
				}
			}
			res = append(res, []int{start, end})
			res = append(res, intervals[i:]...)
			return res
		}

		if i < len(intervals)-1 && nx > x && ny < intervals[i+1][0] {
			res = append(res, intervals[i])
			res = append(res, newInterval)
			res = append(res, intervals[i+1:]...)

			return res
		}

		//default
		res = append(res, [][]int{intervals[i]}...)

	}

	return res
}
