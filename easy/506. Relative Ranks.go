package easy

import (
	"container/heap"
	"fmt"
)

func findRelativeRanks(score []int) []string {
	res := []string{}

	mHeap := &maxHeap{}
	heap.Init(mHeap)

	for i, s := range score {
		heap.Push(mHeap, &hVal{
			rank: s,
			idx:  i,
		})
	}

	nameMap := map[int]string{
		1: "Gold Medal",
		2: "Silver Medal",
		3: "Bronze Medal",
	}

	place := 1
	for mHeap.Len() > 0 {
		val := heap.Pop(mHeap).(hVal)

		if v, ok := nameMap[place]; ok {
			res[val.idx] = v
		} else {
			res[val.idx] = fmt.Sprintf("%d", place)
		}

		place++
	}

	return res
}

type hVal struct {
	rank, idx int
}

type maxHeap []hVal

func (m maxHeap) Less(i, j int) bool {
	return m[i].rank > m[j].rank
}

func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m maxHeap) Len() int {
	return len(m)
}

func (m *maxHeap) Push(v interface{}) {
	val := v.(hVal)
	*m = append(*m, val)
}

func (m *maxHeap) Pop() interface{} {
	if len(*m) == 0 {
		return nil
	}

	top := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]

	return top
}
