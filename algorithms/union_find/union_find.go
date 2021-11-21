package union_find

// Time complexity - O(logn)

type UnionFind struct {
	roots []int
	size  []int
	rank  []int
}

func NewUnionFInd(n int) *UnionFind {
	u := &UnionFind{
		roots: make([]int, n),
		size:  make([]int, n),
		rank:  make([]int, n),
	}

	for i := 0; i < n; i++ {
		u.roots[i] = i
	}

	return u
}

func (u *UnionFind) Find(x int) int {
	for x != u.roots[x] {
		u.roots[x] = u.roots[u.roots[x]]
		x = u.roots[x]
	}

	return x
}

func (u *UnionFind) Union(p, q int) {
	x, y := u.Find(p), u.Find(q)
	u.roots[x] = y
}

func (u *UnionFind) UnionBySize(p, q int) {
	x, y := u.Find(p), u.Find(q)
	if x == y {
		return
	}
	if u.size[x] <= u.size[y] {
		u.roots[x] = u.roots[y]
		u.size[y] += u.size[x]
	} else {
		u.roots[y] = u.roots[x]
		u.size[x] += u.size[y]
	}
}

func (u *UnionFind) UnionByRank(p, q int) {
	x, y := u.Find(p), u.Find(q)
	if x == y {
		return
	}
	if u.rank[x] < u.rank[y] {
		u.roots[x] = y
	} else if u.rank[x] > u.rank[y] {
		u.roots[y] = x
	} else {
		u.roots[x] = y
		u.rank[y]++
	}
}
