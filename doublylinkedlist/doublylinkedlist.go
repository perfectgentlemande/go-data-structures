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
func (d *DoublyLinkedList[T]) LeftPush(val T) {
	newNode := &Node[T]{
		Val: val,
	}

	if d.tail == nil {
		d.tail = d.head
	}

	newNode.Next = d.head
	d.head.Previous = newNode
	d.head = newNode
}

func (d *DoublyLinkedList[T]) RightPush(val T) {
	newNode := &Node[T]{
		Val: val,
	}

	if d.tail == nil {
		d.tail = d.head
	}

	d.tail.Next = newNode
	newNode.Previous = d.tail
	d.tail = newNode
}
