package linkedlist

import "fmt"

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
