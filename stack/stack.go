package stack

type Stack[T any] struct {
	stack []T
}

func New[T any](items ...T) Stack[T] {
	stack := Stack[T]{stack: make([]T, 0)}
	for _, item := range items {
		stack.Push(item)
	}

	return stack
}

func (s *Stack[T]) Push(item T) {
	s.stack = append(s.stack, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.stack) == 0 {
		var zero T

		return zero, false
	}

	pop := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	return pop, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s == nil {
		var zero T

		return zero, false
	}

	return s.stack[len(s.stack)-1], true
}

func (s *Stack[T]) Len() int {
	return len(s.stack)
}

func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}
