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
