package algorithms

type UnionFind struct {
	roots, ranks []int
}

func NewUnionFind(size int) *UnionFind {
	roots := make([]int, size)
	ranks := make([]int, size)

	for i := 0; i < size; i++ {
		roots[i] = i
	}

	return &UnionFind{roots, ranks}
}

func (u UnionFind) Find(x int) int {
	for x != u.roots[x] {
		u.roots[x] = u.roots[u.roots[x]]
		x = u.roots[x]
	}

	return x
}

func (u UnionFind) Union(x, y int) {
	r1 := u.Find(x)
	r2 := u.Find(x)

	if r1 == r2 {
		return
	}

	if u.ranks[r1] > u.ranks[r2] {
		u.roots[r2] = r1
	} else if u.ranks[r2] > u.ranks[r1] {
		u.roots[r1] = r2
	} else {
		u.roots[r1] = r2
		u.ranks[r2]++
	}
}
