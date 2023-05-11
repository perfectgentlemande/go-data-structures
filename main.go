package main

import (
	"fmt"

	"github.com/perfectgentlemande/go-data-structures/deque"
)

func main() {
	deque := deque.New[int]()
	fmt.Printf("deque is empty: %v\n", deque.IsEmpty())
	deque.PushFront(1)
	deque.PushFront(2)
	deque.PushFront(3)
	deque.PushFront(4)
	fmt.Printf("deque is empty: %v\n", deque.IsEmpty())
	v, ok := deque.Head()
	fmt.Printf("head v: %v ok: %v\n", v, ok)
	v, ok = deque.Tail()
	fmt.Printf("tail v: %v ok: %v\n", v, ok)
	deque.PushBack(5)
	deque.PushBack(6)
	v, ok = deque.Tail()
	fmt.Printf("tail v: %v ok: %v\n", v, ok)

	v, ok = deque.PopFront()
	fmt.Printf("popped from front v: %v ok: %v\n", v, ok)
	v, ok = deque.PopFront()
	fmt.Printf("popped from front v: %v ok: %v\n", v, ok)
	v, ok = deque.Head()
	fmt.Printf("head v: %v ok: %v\n", v, ok)

	v, ok = deque.PopBack()
	fmt.Printf("popped from back v: %v ok: %v\n", v, ok)

	v, ok = deque.Tail()
	fmt.Printf("tail v: %v ok: %v\n", v, ok)

	v, ok = deque.PopBack()
	fmt.Printf("popped from back v: %v ok: %v\n", v, ok)
	v, ok = deque.PopBack()
	fmt.Printf("popped from back v: %v ok: %v\n", v, ok)
	v, ok = deque.PopBack()
	fmt.Printf("popped from back v: %v ok: %v\n", v, ok)
	v, ok = deque.PopBack()
	fmt.Printf("popped from back v: %v ok: %v\n", v, ok)

	v, ok = deque.Tail()
	fmt.Printf("tail v: %v ok: %v\n", v, ok)
	v, ok = deque.Head()
	fmt.Printf("head v: %v ok: %v\n", v, ok)
}
