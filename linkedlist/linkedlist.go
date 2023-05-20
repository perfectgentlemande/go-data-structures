package linkedlist

import (
	"fmt"
)

type Node[T int | string] struct {
	Val  T
	Next *Node[T]
}

func (n *Node[T]) PrintList() {
	cur := n

	for cur != nil {
		fmt.Println("val: ", cur.Val)
		cur = cur.Next
	}
}

func (n *Node[T]) CopyList() *Node[T] {
	if n == nil {
		return nil
	}
	cur := n

	newRoot := &Node[T]{
		Val: cur.Val,
	}
	newCur := newRoot

	cur = cur.Next
	for cur != nil {
		newCur.Next = &Node[T]{
			Val: cur.Val,
		}

		newCur = newCur.Next
		cur = cur.Next
	}

	return newRoot
}

func (n *Node[T]) InvertList() *Node[T] {
	if n == nil || n.Next == nil {
		return nil
	}

	prev := n
	cur := prev.Next
	prev.Next = nil

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func (n *Node[T]) InvertListWithCopying() *Node[T] {
	if n == nil || n.Next == nil {
		return nil
	}

	prev := &Node[T]{
		Val: n.Val,
	}

	cur := prev.Next
	for cur != nil {
		prev = &Node[T]{
			Val:  cur.Val,
			Next: prev,
		}

		cur = cur.Next
	}

	return n
}
