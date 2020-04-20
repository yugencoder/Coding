package main

import (
	"fmt"

	"../LC/medium"
)

func main() {
	type TestCases struct {
		s string
		k int
	}

	testCases := []TestCases{{"cxayrgpcctwlfupgzirmazszgfiusitvzsnngmivctprcotcuutfxdpbrdlqukhxkrmpwqqwdxxrptaftpnilfzcmknqljgbfkzuhksxzplpoozablefndimqnffrqfwgaixsovmmilicjwhppikryerkdidupvzdmoejzczkbdpfqkgpbxcrxphhnxfazovxbvaxyxhgqxcxirjsryqxtreptusvupsstylpjniezyfokbowpbgxbtsemzsvqzkbrhkvzyogkuztgfmrprz", 5}}

	for _, tc := range testCases {
		fmt.Printf("Input : %v\n", tc)
		fmt.Printf("Output: %v\n", problem.CanConstruct(tc.s, tc.k))
	}
}
