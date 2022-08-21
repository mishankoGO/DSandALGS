package stacks

// Stack class
type Stack struct {
	Items []int
}

// Push method
func (s *Stack) Push(i int) {
	s.Items = append(s.Items, i)
}

// Pop method
func (s *Stack) Pop() int {
	l := len(s.Items) - 1
	item := s.Items[l]
	s.Items = s.Items[:l]
	return item
}

// Peek method
func (s *Stack) Peek() int {
	l := len(s.Items) - 1
	item := s.Items[l]
	return item
}
