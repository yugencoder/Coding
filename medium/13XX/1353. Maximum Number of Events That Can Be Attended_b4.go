package t13XX

//
//type Item struct {
//	index      int
//	value      int
//	currentDay int
//	endDay     int
//}
//
//type PriorityQueue []*Item
//
//func (pq PriorityQueue) Len() int { return len(pq) }
//
//func (pq PriorityQueue) Less(i, j int) bool {
//	return pq[i].endDay < pq[j].endDay
//}
//
//func (pq *PriorityQueue) Pop() interface{} {
//	old := *pq
//	n := len(*pq)
//	item := old[n-1]
//	old[n-1] = nil  // avoid memory leak
//	item.index = -1 // for safety
//	*pq = old[0 : n-1]
//	return item
//}
//
//func (pq *PriorityQueue) Push(x interface{}) {
//	n := len(*pq)
//	item := x.(*Item)
//	item.index = n
//	*pq = append(*pq, item)
//}
//
//func (pq PriorityQueue) Swap(i, j int) {
//	pq[i], pq[j] = pq[j], pq[i]
//	pq[i].index = i
//	pq[j].index = j
//}
//
//func maxEvents(events [][]int) int {
//	pq := PriorityQueue{}
//	eventCount := 0
//
//	// sort all events
//	sort.Slice(events, func(i, j int) bool {
//		if events[i][0] == events[j][0] {
//			return events[i][1] < events[j][1]
//		}
//		return events[i][0] < events[j][0]
//	})
//
//	fmt.Printf("All items: %v\n", events)
//
//	j := 0
//	for i := 1; i <= 100000; i++ {
//
//		for ; j < len(events) && events[j][0] == i; j++ {
//			heap.Push(&pq, &Item{
//				endDay: events[j][1],
//				value:  j + 1,
//			})
//
//		}
//
//		for ; len(pq) > 0; {
//			itIf := heap.Pop(&pq)
//			item := itIf.(*Item)
//			if i > item.endDay {
//				continue
//			} else {
//				fmt.Printf("Picked item: %v for day: %v\n", item.value, i)
//				eventCount++
//				break
//			}
//		}
//	}
//
//	return eventCount
//}
