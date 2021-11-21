package problem

import "container/heap"

//1648. Sell Diminishing-Valued Colored Balls

func maxProfit(inventory []int, orders int) int {
	maxHeap := &maxHeap{}
	heap.Init(maxHeap)

	for _, inv := range inventory {
		heap.Push(maxHeap, inv)
	}

	res := 0
	for ; orders > 0; {
		if maxHeap.Len() == 0 {
			return res
		}

		maxVal := heap.Pop(maxHeap).(int)
		nextVal := 0

		if maxHeap.Len() > 0 {
			nextVal = heap.Pop(maxHeap).(int)
		}

		if maxVal > nextVal {
			nextMin := min(orders,maxVal)
			res += maxVal*(maxVal+1)/2 - nextVal*(nextVal+1)/2
			orders -= maxVal - nextVal
		} else {
			res += maxVal
			orders--
		}

		if maxVal > 1 {
			heap.Push(maxHeap, maxVal-1)
		}
		heap.Push(maxHeap, nextVal)
	}

	return res
}

type maxHeap []int

func (m maxHeap) Len() int {
	return len(m)
}

func (m maxHeap) Less(a, b int) bool {
	return m[a] > m[b]
}

func (m maxHeap) Swap(a, b int) {
	m[a], m[b] = m[b], m[a]
}

func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *maxHeap) Pop() interface{} {
	n := len(*m)
	x := (*m)[n-1]
	*m = (*m)[0 : n-1]
	return x
}
