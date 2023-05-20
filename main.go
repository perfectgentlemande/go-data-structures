package main

import "github.com/perfectgentlemande/go-data-structures/linkedlist"

func main() {
	node := linkedlist.Node[int]{
		Val: 1,
		Next: &linkedlist.Node[int]{
			Val: 2,
			Next: &linkedlist.Node[int]{
				Val: 3,
				Next: &linkedlist.Node[int]{
					Val: 4,
					Next: &linkedlist.Node[int]{
						Val: 5,
					},
				},
			},
		},
	}

	node.PrintList()

	newNode := node.CopyList()
	newNode.PrintList()
	newNode.Val = 6
	node.PrintList()
	newNode.PrintList()
}
