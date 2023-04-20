package doublylinkedlist

import "fmt"

type Node[T int | string] struct {
	Val      T
	Next     *Node[T]
	Previous *Node[T]
}

type DoublyLinkedList[T int | string] struct {
	head *Node[T]
	tail *Node[T]
}

func New[T int | string]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (d *DoublyLinkedList[T]) PrintList() {
	cur := d.head

	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}
func (d *DoublyLinkedList[T]) PrintListBackwards() {
	cur := d.tail

	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Previous
	}
}
