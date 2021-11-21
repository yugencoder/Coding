package _7XX

import (
	"container/heap"
	"math"
)

type Conn struct {
	Weight int
	Vertex int
}

type MinHeap []Conn

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].Weight < h[j].Weight }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Conn)) }
func (h *MinHeap) Pop() (v interface{}) {
	*h, v = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}

func countRestrictedPaths(n int, edges [][]int) int {
	graph := make(map[int][]Conn, n+1)
	for i := 1; i <= n; i++ {
		graph[i] = make([]Conn, 0)
	}

	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		graph[u] = append(graph[u], Conn{w, v})
		graph[v] = append(graph[v], Conn{w, u})
	}

	st := findShortestPath(n, graph)

	memo = make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}

	return dfsRestrictedPaths(1, n, graph, st)
}

var memo []int

func dfsRestrictedPaths(u int, n int, graph map[int][]Conn, sd []int) int {
	if memo[u] >= 0 {
		return memo[u]
	}

	currMemo := 0
	mod := 1000000007

	for _, conn := range graph[u] {
		v := conn.Vertex

		if sd[u] > sd[v] {
			if v == n {
				currMemo = (currMemo + 1) % mod
			} else {
				currMemo = (currMemo + dfsRestrictedPaths(v, n, graph, sd)) % mod
			}
		}
	}

	memo[u] = currMemo

	return currMemo
}

func findShortestPath(start int, graph map[int][]Conn) []int {
	n := start

	dist := make([]int, len(graph)+1)
	for i := range dist {
		dist[i] = math.MaxInt32
	}

	dist[n] = 0

	h := &MinHeap{}
	heap.Init(h)
	h.Push(Conn{0, n})

	for h.Len() > 0 {
		u := heap.Pop(h).(Conn)
		if u.Weight > dist[u.Vertex] {
			continue
		}
		for _, v := range graph[u.Vertex] {
			if dist[u.Vertex]+v.Weight < dist[v.Vertex] {
				dist[v.Vertex] = dist[u.Vertex] + v.Weight
				heap.Push(h, Conn{dist[v.Vertex], v.Vertex})
			}
		}
	}

	return dist
}
