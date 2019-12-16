package stack

// Stack is a FILO datastruct
type Stack struct {
	data []int
	len  int
}

// New return a stack
func New() *Stack {
	return &Stack{make([]int, 0), 0}
}

// Push push a data
func (s *Stack) Push(v int) {
	s.data = append(s.data, v)
	s.len++
}

// Pop pop a data
func (s *Stack) Pop() (int, bool) {
	if s.len > 0 {
		ret := s.data[0]
		copy(s.data, s.data[1:])
		s.len--
		s.data = s.data[:s.len]
		return ret, true
	}
	return 0, false
}

// Len return stack's len
func (s *Stack) Len() int {
	return s.len
}
