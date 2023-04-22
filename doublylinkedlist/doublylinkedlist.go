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

	if d.head == nil {
		d.head = newNode
		d.tail = d.head
		return
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
		d.tail = newNode
		d.head = d.tail
		return
	}

	d.tail.Next = newNode
	newNode.Previous = d.tail
	d.tail = newNode
}
func (d *DoublyLinkedList[T]) LeftTop() (T, bool) {
	if d == nil || d.head == nil {
		return Node[T]{}.Val, false
	}

	return d.head.Val, true
}

func (d *DoublyLinkedList[T]) RightTop() (T, bool) {
	if d == nil || d.tail == nil {
		return Node[T]{}.Val, false
	}

	return d.tail.Val, true
}

func (d *DoublyLinkedList[T]) LeftPop() (T, bool) {
	if d == nil || d.head == nil {
		return Node[T]{}.Val, false
	}

	cur := d.head.Val
	if d.tail == d.head {
		d.tail = nil
	}

	d.head = d.head.Next

	if d.head != nil {
		d.head.Previous = nil
	}

	return cur, true
}

func (d *DoublyLinkedList[T]) RightPop() (T, bool) {
	if d == nil || d.tail == nil {
		return Node[T]{}.Val, false
	}

	cur := d.tail.Val
	if d.tail == d.head {
		d.head = nil
	}

	d.tail = d.tail.Previous

	if d.tail != nil {
		d.tail.Next = nil
	}

	return cur, true
}
