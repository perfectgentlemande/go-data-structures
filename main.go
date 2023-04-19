package main

import (
	"fmt"

	"github.com/perfectgentlemande/go-data-structures/linkedlist"
	"github.com/perfectgentlemande/go-data-structures/queue"
	"github.com/perfectgentlemande/go-data-structures/stack"
)

func main() {
	root := &linkedlist.Node[int]{
		Val: 1,
		Next: &linkedlist.Node[int]{
			Val: 2,
			Next: &linkedlist.Node[int]{
				Val: 3,
			},
		},
	}

	root.PrintList()

	stack := stack.New[int]()
	v, ok := stack.Top()
	fmt.Printf("stack is empty: %v\n", stack.IsEmpty())
	fmt.Printf("v: %v ok: %v\n", v, ok)
	stack.Push(1)
	v, ok = stack.Top()
	fmt.Printf("v: %v ok: %v\n", v, ok)
	stack.Push(2)
	v, ok = stack.Top()
	fmt.Printf("v: %v ok: %v\n", v, ok)
	stack.Push(3)
	v, ok = stack.Top()
	fmt.Printf("v: %v ok: %v\n", v, ok)
	fmt.Printf("stack is empty: %v\n", stack.IsEmpty())
	v, ok = stack.Pop()
	fmt.Printf("popped v: %v ok: %v\n", v, ok)
	v, ok = stack.Top()
	fmt.Printf("after pop v: %v ok: %v\n", v, ok)
	v, ok = stack.Pop()
	fmt.Printf("popped v: %v ok: %v\n", v, ok)
	v, ok = stack.Top()
	fmt.Printf("after pop v: %v ok: %v\n", v, ok)
	v, ok = stack.Pop()
	fmt.Printf("popped v: %v ok: %v\n", v, ok)
	v, ok = stack.Top()
	fmt.Printf("after pop v: %v ok: %v\n", v, ok)
	v, ok = stack.Pop()
	fmt.Printf("popped v: %v ok: %v\n", v, ok)
	v, ok = stack.Top()
	fmt.Printf("after pop v: %v ok: %v\n", v, ok)
	fmt.Printf("stack is empty: %v\n", stack.IsEmpty())

	queue := queue.New[int]()
	fmt.Printf("queue is empty: %v\n", queue.IsEmpty())
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)
	fmt.Printf("queue is empty: %v\n", queue.IsEmpty())
	v, ok = queue.Remove()
	fmt.Printf("after removal v: %v ok: %v\n", v, ok)
	v, ok = queue.Remove()
	fmt.Printf("after removal v: %v ok: %v\n", v, ok)
	v, ok = queue.Remove()
	fmt.Printf("after removal v: %v ok: %v\n", v, ok)
	fmt.Printf("queue is empty: %v\n", queue.IsEmpty())
}
