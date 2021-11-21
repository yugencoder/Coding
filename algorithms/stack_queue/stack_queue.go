package stack_queue

// A Tree is a simple binary heap
type Tree struct {
	root *TNode
}

type TNode struct {
	key   int
	value int
	level int
	left  *TNode
	right *TNode
}
