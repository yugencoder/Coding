package X

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	res := [][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] <= intervals[j][0] {
			return true
		}

		return false
	})

	start := intervals[0][0]
	end := intervals[0][1]

	if len(intervals) == 0 {
		return [][]int{{start, end}}
	}

	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]

		if max(end, interval[1])-min(start, interval[0]) <= (interval[1] - interval[0] + end - start) {
			start = min(start, interval[0])
			end = max(end, interval[1])
		} else {
			res = append(res, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]
		}

		if i == len(intervals)-1 {
			res = append(res, []int{start, end})
		}
	}

	return res
}
