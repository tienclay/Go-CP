package ds

// Stack
type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(x T) {
	s.data = append(s.data, x)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		return *new(T), false
	}
	x := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return x, true
}
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.data) == 0 {
		return *new(T), false
	}
	return s.data[len(s.data)-1], true
}
func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) Empty() bool {
	return len(s.data) == 0
}
