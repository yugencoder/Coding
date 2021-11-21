package stack_queue

type Stack []TNode

func NewStack() Stack {
	node := make([]TNode, 0)
	return node
}

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(node *TNode) {
	*s = append(*s, *node) // Simply append the new value to the end of the stack
}

// Pop Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() *TNode {
	if s.IsEmpty() {
		return nil
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return &element
	}
}
