package stack_queue

// BFS traverses the heap in the level order
func (t *Tree) BFS() {
	if t.root != nil {
		queue := New()
		queue.EnQueue(t.root)
		for {
			if queue.IsEmpty() {
				break
			}
			n := queue.DeQueue()
			if n.left != nil {
				queue.EnQueue(n.left)
			}
			if n.right != nil {
				queue.EnQueue(n.right)
			}
		}
	}
	return
}

// DFS traverse the heap in pre-order
func (t *Tree) DFS() {

	if t.root != nil {
		stack := NewStack()
		stack.Push(t.root)

		for {
			if stack.IsEmpty() {
				break
			}
			n := stack.Pop()
			if n.right != nil {
				stack.Push(n.right)
			}
			if n.left != nil {
				stack.Push(n.left)
			}
		}
	}
}
