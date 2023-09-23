package service

type Stack[T any] []T

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Size() int {
	return len(*s)
}

func (s *Stack[T]) Push(element T) {
	*s = append(*s, element)
}

func (s *Stack[T]) Top() *T {
	if s.IsEmpty() {
		return nil
	}
	return &(*s)[s.Size()-1]
}

func (s *Stack[T]) Pop() *T {
	if s.IsEmpty() {
		return nil
	}
	n := s.Size() - 1
	item := (*s)[n]
	*s = (*s)[:n]
	return &item
}
