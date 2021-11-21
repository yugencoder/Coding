package _7XX

import "container/heap"

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() (v interface{}) {
	*h, v = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}

func maximumScore(a int, b int, c int) int {
	h := &MaxHeap{}
	heap.Init(h)

	heap.Push(h, a)
	heap.Push(h, b)
	heap.Push(h, c)

	res := 0
	for ; ; {
		v1 := heap.Pop(h).(int)
		v2 := heap.Pop(h).(int)

		if v1 <= 0 || v2 <= 0 {
			return res
		}

		heap.Push(h, v1-1)
		heap.Push(h, v2-1)
		res++
	}

	return res
}
