package main

import (
	"fmt"

	"../LC/medium"
)

func main() {
	type TestCases struct {
		nums []int
	}

	testCases := []TestCases{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}}

	for _, tc := range testCases {
		fmt.Printf("Input : %v\n", tc)
		fmt.Printf("Output: %v\n", problem.SumFourDivisors(tc.nums))
	}
}
