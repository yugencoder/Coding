package t13XX

import (
	"container/heap"
	"sort"
)

type PriorityQueue []int

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i] < pq[j] }

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func maxEvents(events [][]int) int {
	pq := &PriorityQueue{}
	eventCount := 0

	// sort all events
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	j := 0
	for i := 1; i <= 100000; i++ {
		for ; j < len(events) && events[j][0] == i; j++ {
			heap.Push(pq, events[j][1])
		}

		for ; pq.Len() > 0; {
			item := heap.Pop(pq)
			if i > item.(int) {
				continue
			}
			eventCount++
			break
		}
	}

	return eventCount
}
