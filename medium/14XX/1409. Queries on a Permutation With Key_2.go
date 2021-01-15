package _14XX

/*
Given the array queries of positive integers between 1 and m, you have to process all queries[key] (from key=0 to key=queries.length-1) according to the following rules:

In the beginning, you have the permutation P=[1,2,3,...,m].
For the current key, find the position of queries[key] in the permutation P (indexing from 0) and then move this at the beginning of the permutation P. Notice that the position of queries[key] in P is the result for queries[key].
Return an array containing the result for the given queries.



Example 1:

Input: queries = [3,1,2,1], m = 5
Output: [2,1,2,1]
Explanation: The queries are processed as follow:
For key=0: queries[key]=3, P=[1,2,3,4,5], position of 3 in P is 2, then we move 3 to the beginning of P resulting in P=[3,1,2,4,5].
For key=1: queries[key]=1, P=[3,1,2,4,5], position of 1 in P is 1, then we move 1 to the beginning of P resulting in P=[1,3,2,4,5].
For key=2: queries[key]=2, P=[1,3,2,4,5], position of 2 in P is 2, then we move 2 to the beginning of P resulting in P=[2,1,3,4,5].
For key=3: queries[key]=1, P=[2,1,3,4,5], position of 1 in P is 1, then we move 1 to the beginning of P resulting in P=[1,2,3,4,5].
Therefore, the array containing the result is [2,1,2,1].
Example 2:

Input: queries = [4,1,2,2], m = 4
Output: [3,1,2,0]
Example 3:

Input: queries = [7,5,5,8,3], m = 8
Output: [6,5,0,7,5]
*/

type Fenwick struct {
	arr []int
}

func NewFenwick(size int) *Fenwick {
	return &Fenwick{
		arr: make([]int, size),
	}
}

func (f *Fenwick) Update(n int, val int) {
	for ; n < len(f.arr); n += (n & (-n)) {
		f.arr[n] += val
	}
}

func (f *Fenwick) Query(n int) int {
	sum := 0

	for ; n > 0; n = n - (n & -n) {
		sum += f.arr[n]
	}

	return sum
}

func ProcessQueries_2(queries []int, m int) []int {
	bit := NewFenwick(2*m + 1)
	mapping := map[int]int{}
	res := make([]int, len(queries))

	for i := 1; i <= m; i++ {
		mapping[i] = m + i
		bit.Update(m+i, 1)
	}

	for i, q := range queries {
		res[i] = bit.Query(mapping[q]) - 1

		bit.Update(mapping[q], -1)
		bit.Update(m, 1)
		mapping[q] = m
		m--
	}

	return res
}
