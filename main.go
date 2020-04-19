package main

import (
	"fmt"

	"../LC/medium"
)

func main() {
	type TestCases struct {
		queries []int
		m       int
	}
	testCases := []TestCases{{[]int{3, 1, 2, 1}, 5}}

	for _, tc := range testCases {
		fmt.Println("Input : %v", tc)
		fmt.Println(problem.ProcessQueries_2(tc.queries, tc.m))
	}
}
