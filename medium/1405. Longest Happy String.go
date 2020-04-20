package problem

import "container/heap"

/*
A string is called happy if it does not have any of the strings 'aaa', 'bbb' or 'ccc' as a substring.

Given three integers a, b and c, return any string s, which satisfies following conditions:

s is happy and longest possible.
s contains at most a occurrences of the letter 'a', at most b occurrences of the letter 'b' and at most c occurrences of the letter 'c'.
s will only contain 'a', 'b' and 'c' letters.
If there is no such string s return the empty string "".



Example 1:

Input: a = 1, b = 1, c = 7
Output: "ccaccbcc"
Explanation: "ccbccacc" would also be a correct answer.
Example 2:

Input: a = 2, b = 2, c = 1
Output: "aabbc"
Example 3:

Input: a = 7, b = 1, c = 0
Output: "aabaa"
Explanation: It's the only correct answer in this case.
*/

type pair struct {
	char  byte
	count int
}

type maxHeap []pair

func (m maxHeap) Len() int {
	return len(m)
}

func (m maxHeap) Less(a, b int) bool {
	return m[a].count > m[b].count
}

func (m maxHeap) Swap(a, b int) {
	m[a], m[b] = m[b], m[a]
}

func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(pair))
}

func (m *maxHeap) Pop() interface{} {
	n := len(*m)
	x := (*m)[n-1]
	*m = (*m)[0 : n-1]
	return x
}

func LongestDiverseString(a int, b int, c int) string {
	mHeap := &maxHeap{}
	heap.Init(mHeap)
	res := []byte{}

	if a > 0 {
		heap.Push(mHeap, pair{'a', a})
	}

	if b > 0 {
		heap.Push(mHeap, pair{'b', b})
	}

	if c > 0 {
		heap.Push(mHeap, pair{'c', c})
	}

	for mHeap.Len() > 0 {
		firstP := heap.Pop(mHeap).(pair)
		if len(res) == 0 || firstP.char != res[len(res)-1] {
			res = append(res, firstP.char)
			firstP.count = firstP.count - 1

			if firstP.count >= 2 {
				res = append(res, firstP.char)
				firstP.count = firstP.count - 1
			}
		}

		if mHeap.Len() == 0 {
			break
		}

		secondP := heap.Pop(mHeap).(pair)
		res = append(res, secondP.char)
		secondP.count = secondP.count - 1

		if secondP.count >= 2 && secondP.count > firstP.count {
			res = append(res, firstP.char)
			firstP.count = firstP.count - 1
		}

		if firstP.count > 0 {
			heap.Push(mHeap, firstP)
		}

		if secondP.count > 0 {
			heap.Push(mHeap, secondP)
		}
	}

	return string(res)
}
