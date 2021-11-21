package _8XX

import "math"

func countPoints(points [][]int, queries [][]int) []int {
	res := make([]int, len(queries))

	for i, q := range queries {
		x := q[0]
		y := q[1]
		r := q[2]

		for _, p := range points {
			if int(math.Abs(float64(p[0]-x)))*int(math.Abs(float64(p[0]-x)))+int(math.Abs(float64(p[1]-y)))*int(math.Abs(float64(p[1]-y))) <= r*r {
				res[i]++
			}
		}
	}

	return res
}
