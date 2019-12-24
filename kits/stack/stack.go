package stack

// Stack is a FILO datastruct
type Stack struct {
	data []interface{}
	len  int
}

// New return a stack
func New() *Stack {
	return &Stack{make([]interface{}, 0), 0}
}

// Push push a data
func (s *Stack) Push(v interface{}) {
	s.data = append(s.data, v)
	s.len++
}

// Pop pop a data
func (s *Stack) Pop() (interface{}, bool) {
	if s.len > 0 {
		s.len--
		ret := s.data[s.len]
		s.data = s.data[:s.len]
		return ret, true
	}
	return 0, false
}

// Len return stack's len
func (s *Stack) Len() int {
	return s.len
}
