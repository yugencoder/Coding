package _7XX

type stack []element

func (s *stack) Pop() element {
	temp := (*s)[0]
	*s = (*s)[1:]

	return temp
}

func (s *stack) Push(e element) {
	n := len(*s)

	*s = append(*s, e)
	copy((*s)[1:], (*s)[0:n])
	(*s)[0] = e

	return
}

func (s *stack) Len() int {
	return len(*s)
}

type element struct {
	i       int
	visited map[int]bool
}
