package main

import (
	"fmt"

	"github.com/perfectgentlemande/go-data-structures/linkedlist"
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
}
