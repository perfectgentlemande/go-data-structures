package deque

import (
	"github.com/perfectgentlemande/go-data-structures/doublylinkedlist"
	"github.com/perfectgentlemande/go-data-structures/linkedlist"
)

type Deque[T int | string] struct {
	dll *doublylinkedlist.DoublyLinkedList[T]
}

func New[T int | string]() *Deque[T] {
	return &Deque[T]{}
}

func (d *Deque[T]) Head() (T, bool) {
	if d == nil || d.dll == nil {
		return linkedlist.Node[T]{}.Val, false
	}

	return d.dll.LeftTop()
}
func (d *Deque[T]) Tail() (T, bool) {
	if d == nil || d.dll == nil {
		return linkedlist.Node[T]{}.Val, false
	}

	return d.dll.RightTop()
}

func (d *Deque[T]) PushBack(val T) {
	d.dll.RightPush(val)
}
func (d *Deque[T]) PushFront(val T) {
	d.dll.LeftPush(val)
}

func (d *Deque[T]) PopBack() (T, bool) {
	if d == nil || d.dll == nil {
		return linkedlist.Node[T]{}.Val, false
	}

	return d.dll.RightPop()
}
func (d *Deque[T]) PopFront() (T, bool) {
	if d == nil || d.dll == nil {
		return linkedlist.Node[T]{}.Val, false
	}

	return d.dll.LeftPop()
}

func (d *Deque[T]) IsEmpty() bool {
	return d == nil || d.dll == nil || d.dll.IsEmpty()
}
