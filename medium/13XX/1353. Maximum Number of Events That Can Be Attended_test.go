package t13XX

import (
	"fmt"
	"testing"
)

func TestMaxEvents(t *testing.T) {
	allEvents := [][]int{
		//{1, 2}, {2, 3}, {3, 4}, {1, 2},
		//{1,4},{4,4},{2,2},{3,4},{1,1},
		{1, 2}, {1, 2}, {1, 6}, {1, 2}, {1, 2},
		//{1, 2}, {2,3},{3,4},
	}

	fmt.Println("\nresult",maxEvents(allEvents))
	//
	//pq := &PriorityQueue{}
	//heap.Push(pq, &Item{
	//	index: 1,
	//	value: 1,
	//	endDay:  2,
	//})
	//heap.Push(pq, &Item{
	//	index: 2,
	//	value: 2,
	//	endDay:  3,
	//})
	//heap.Push(pq, &Item{
	//	index: 3,
	//	value: 4,
	//	endDay:  5,
	//})
	//heap.Push(pq, &Item{
	//	index: 0,
	//	value: 3,
	//	endDay:  4,
	//})
	//
	//for ; len(*pq) > 0; {
	//	item := heap.Pop(pq)
	//	it := item.(*Item)
	//	fmt.Println(it.endDay)
	//}

}
