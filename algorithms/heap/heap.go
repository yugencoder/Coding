package heap

type Element interface {
	LessThan(e Element) bool
}

type Heap struct {
	heap  []Element
	len   int
	isMin bool
}

func New(isMin bool) *Heap {
	return &Heap{heap: make([]Element, 0), isMin: isMin}
}

func (h *Heap) IsEmpty() bool {
	return len(h.heap) == 0
}

func (h *Heap) Len() int {
	return h.len
}

func (h *Heap) Insert(e Element) {

	h.heap = append(h.heap, e)
	h.len++
	h.percolateUp(h.len - 1)
}

func (h *Heap) Extract() (e Element) {
	if h.len == 0 {
		panic("Empty heap, cannot Extract.")
	}

	e, h.heap[0] = h.heap[0], h.heap[h.len-1]
	h.heap = h.heap[:h.len-1]
	h.len--
	h.percolateDown(0)
	return
}

func (h *Heap) Peek() (e Element) {
	if h.len == 0 {
		panic("Empty heap, cannot Peek.")
	}
	return h.heap[0]
}

func (h *Heap) percolateUp(index int) {
	needUp, parent := index, (index-1)>>1
	for needUp > 0 && h.less(h.heap[needUp], h.heap[parent]) {
		h.heap[parent], h.heap[needUp] = h.heap[needUp], h.heap[parent]
		needUp, parent = parent, parent>>1
	}
}

func (h *Heap) percolateDown(index int) {
	for needDown, child := index, index<<1+1; needDown < h.len && child < h.len; needDown, child = child, child<<1+1 {
		// find min(leftChild, rightChild)
		if child+1 < h.len && h.less(h.heap[child+1], h.heap[child]) {
			child++
		}
		if h.less(h.heap[needDown], h.heap[child]) {
			break
		}
		h.heap[needDown], h.heap[child] = h.heap[child], h.heap[needDown]
	}
}

func (h *Heap) less(a, b Element) bool {
	if h.isMin {
		return a.LessThan(b)
	}
	return b.LessThan(a)
}
