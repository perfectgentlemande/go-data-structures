package main

import "github.com/perfectgentlemande/go-data-structures/linkedlist"

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
}
