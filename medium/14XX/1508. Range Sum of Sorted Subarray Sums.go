package _14XX

import "container/heap"

type item struct {
	idx int
	val int
}

type pq []*item

func (p pq) Len() int {
	return len(p)
}

func (p pq) Less(i, j int) bool {
	return p[i].val < p[j].val
}

// Swap swaps the elements with indexes key and j.
func (p pq) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *pq) Push(x interface{}) {
	v := x.(*item)
	*p = append(*p, v)
}

func (p *pq) Pop() interface{} {
	v := (*p)[p.Len()-1]
	*p = (*p)[:p.Len()-1]
	return v
}

func rangeSum(nums []int, n int, left int, right int) int {
	mod := 1000000007
	res := 0
	c := 0
	q := &pq{}

	for i, n := range nums {
		q.Push(&item{
			idx: i,
			val: n,
		})
	}
	heap.Init(q)
	for ; q.Len() > 0; {
		top := heap.Pop(q).(*item)
		c++
		if c >= left {
			res = (res + top.val) % mod
		}

		if c == right {
			break
		}
		if top.idx < len(nums)-1 {
			heap.Push(q, &item{
				idx: top.idx + 1,
				val: (top.val + nums[top.idx+1]) % mod,
			})
		}
	}

	return res
}
