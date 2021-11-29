package golang

import (
	"container/heap"
	"fmt"
)

type heapData struct {
	extraData int
	key       int
}
type maxHeap []*heapData

func (h maxHeap) Less(i, j int) bool {
	return h[i].key > h[j].key

}
func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(a interface{}) {
	*h = append(*h, a.(*heapData))
}

func (h *maxHeap) Pop() interface{} {
	if len(*h) == 0 {
		return nil
	}

	temp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return temp
}

func usePQ(nums []int) {

	var mHeap maxHeap
	mHeap = []*heapData{
		{
			extraData: 2,
			key:       2,
		},
		{
			extraData: 4,
			key:       4,
		},
	}

	heap.Init(&mHeap)
	heap.Push(&mHeap,
		&heapData{
			extraData: 4,
			key:       4,
		})

	ans := heap.Pop(&mHeap)
	fmt.Println(ans.(*heapData))
}
