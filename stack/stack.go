package stack

import "github.com/perfectgentlemande/go-data-structures/linkedlist"

type Stack[T int | string] struct {
	root *linkedlist.Node[T]
}

func (s *Stack[T]) Top() (T, bool) {
	if s == nil {
		return linkedlist.Node[T]{}.Val, false
	}

	return s.root.Val, true
}
