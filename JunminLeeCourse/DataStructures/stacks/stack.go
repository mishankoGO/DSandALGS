package stacks

// Stack class
type Stack struct {
	items []int
}

// Push method
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop method
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	item := s.items[l]
	s.items = s.items[:l]
	return item
}
