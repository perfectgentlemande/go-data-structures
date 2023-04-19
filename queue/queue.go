package queue

import "github.com/perfectgentlemande/go-data-structures/linkedlist"

type Queue[T int | string] struct {
	head *linkedlist.Node[T]
	tail *linkedlist.Node[T]
}

func New[T int | string]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Head() (T, bool) {
	if q == nil || q.head == nil {
		return linkedlist.Node[T]{}.Val, false
	}

	return q.head.Val, true
}

func (q *Queue[T]) Add(val T) {
	newTail := &linkedlist.Node[T]{
		Val: val,
	}

	if q.head == nil {
		q.head = newTail
	}

	if q.tail != nil {
		q.tail.Next = newTail
	}

	q.tail = newTail
}

func (q *Queue[T]) Remove() (T, bool) {
	if q == nil || q.head == nil {
		return linkedlist.Node[T]{}.Val, false
	}

	cur := q.head
	if q.tail == q.head {
		q.tail = nil
	}
	q.head = q.head.Next

	return cur.Val, true
}

func (q *Queue[T]) IsEmpty() bool {
	return q == nil || q.head == nil
}
