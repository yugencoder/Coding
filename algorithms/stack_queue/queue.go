package stack_queue

type Queue struct {
	queue []TNode
}

// New Constructor initialize your data structure here. Set the size of the queue to be k.
func New() Queue {
	return Queue{
		queue: make([]TNode, 0),
	}
}

// EnQueue insert an element into the circular queue. Return true if the operation is successful.
func (m *Queue) EnQueue(node *TNode) bool {
	m.queue = append(m.queue, *node)
	return true
}

// DeQueue delete an element from the circular queue. Return true if the operation is successful.
func (m *Queue) DeQueue() *TNode {
	if len(m.queue) == 0 {
		return nil
	}
	node := m.queue[0]
	m.queue = m.queue[1:]
	return &node
}

// Front get the front item from the queue.
func (m *Queue) Front() *TNode {
	if len(m.queue) == 0 {
		return nil
	}
	return &m.queue[0]
}

// Rear get the last item from the queue. */
func (m *Queue) Rear() *TNode {
	if len(m.queue) == 0 {
		return nil
	}
	return &m.queue[len(m.queue)-1]
}

// IsEmpty checks whether the circular queue is empty or not. */
func (m *Queue) IsEmpty() bool {
	return len(m.queue) == 0
}
