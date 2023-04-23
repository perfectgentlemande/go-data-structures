package main

import (
	"fmt"

	"github.com/perfectgentlemande/go-data-structures/doublylinkedlist"
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

	dll := doublylinkedlist.New[int]()
	dll.LeftPush(3)
	dll.LeftPush(2)
	dll.LeftPush(1)
	fmt.Println("print")
	dll.PrintList()
	fmt.Println("print backwards")
	dll.PrintListBackwards()
	dll.RightPush(4)
	dll.RightPush(5)
	dll.RightPush(6)
	fmt.Println("print")
	dll.PrintList()
	fmt.Println("print backwards")
	dll.PrintListBackwards()
	v, ok = dll.LeftTop()
	fmt.Printf("left top v: %v ok: %v\n", v, ok)
	v, ok = dll.RightTop()
	fmt.Printf("right top v: %v ok: %v\n", v, ok)

	v, ok = dll.LeftPop()
	fmt.Printf("left pop v: %v ok: %v\n", v, ok)
	v, ok = dll.LeftPop()
	fmt.Printf("left pop v: %v ok: %v\n", v, ok)
	v, ok = dll.LeftPop()
	fmt.Printf("left pop v: %v ok: %v\n", v, ok)
	dll.PrintList()
	dll.PrintListBackwards()
	v, ok = dll.LeftPop()
	fmt.Printf("left pop v: %v ok: %v\n", v, ok)
	v, ok = dll.LeftPop()
	fmt.Printf("left pop v: %v ok: %v\n", v, ok)
	v, ok = dll.LeftPop()
	fmt.Printf("left pop v: %v ok: %v\n", v, ok)
	v, ok = dll.LeftPop()
	fmt.Printf("left pop v: %v ok: %v\n", v, ok)

	dll.LeftPush(3)
	dll.LeftPush(2)
	dll.LeftPush(1)
	dll.RightPush(4)
	dll.RightPush(5)
	dll.RightPush(6)
	dll.PrintList()
	dll.PrintListBackwards()

	v, ok = dll.RightPop()
	fmt.Printf("right pop v: %v ok: %v\n", v, ok)
	v, ok = dll.RightPop()
	fmt.Printf("rigth pop v: %v ok: %v\n", v, ok)
	v, ok = dll.RightPop()
	fmt.Printf("right pop v: %v ok: %v\n", v, ok)
	dll.PrintList()
	dll.PrintListBackwards()

	dll.RightPush(4)
	dll.RightPush(5)
	dll.RightPush(6)

	fmt.Println("insert before:")
	dll.PrintList()

	node := dll.ExtractIthNode(3)
	fmt.Printf("extracted %vth node: %v\n", 3, node.Val)

	dll.InsertAfter(7, node)
	dll.PrintList()
	dll.PrintListBackwards()

	node = dll.ExtractIthNode(6)
	fmt.Printf("extracted %vth node: %v\n", 6, node.Val)
	dll.InsertAfter(8, node)

	dll.PrintList()
	dll.PrintListBackwards()
}
