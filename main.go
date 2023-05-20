package main

import (
	"fmt"

	"github.com/perfectgentlemande/go-data-structures/linkedlist"
)

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

	fmt.Println("before inversion")
	newNode.PrintList()
	newNode = newNode.InvertList()
	fmt.Println("after inversion")
	newNode.PrintList()

	newNode = newNode.InvertListWithCopying()
	newNode.PrintList()
}
