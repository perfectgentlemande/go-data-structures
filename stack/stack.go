package stack

import "github.com/perfectgentlemande/go-data-structures/linkedlist"

type Stack[T int | string] struct {
	root *linkedlist.Node[T]
}

func New[T int | string]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Top() (T, bool) {
	if s == nil || s.root == nil {
		return linkedlist.Node[T]{}.Val, false
	}

	return s.root.Val, true
}

func (s *Stack[T]) Push(val T) {
	newNode := &linkedlist.Node[T]{
		Val:  val,
		Next: s.root,
	}

	s.root = newNode
}

func (s *Stack[T]) Pop() (T, bool) {
	if s == nil {
		return linkedlist.Node[T]{}.Val, false
	}

	cur, ok := s.Top()
	if !ok {
		return cur, ok
	}

	s.root = s.root.Next
	return cur, ok
}

func (s *Stack[T]) IsEmpty() bool {
	return s == nil || s.root == nil
}
