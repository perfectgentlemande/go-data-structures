package main

import (
	"fmt"

	"github.com/perfectgentlemande/go-data-structures/binarytree"
)

func main() {
	tree := binarytree.Node[int]{
		Val: 1,
		Left: &binarytree.Node[int]{
			Val: 2,
			Left: &binarytree.Node[int]{
				Val: 4,
			},
			Right: &binarytree.Node[int]{
				Val: 5,
			},
		},
		Right: &binarytree.Node[int]{
			Val: 3,
			Left: &binarytree.Node[int]{
				Val: 6,
			},
			Right: &binarytree.Node[int]{
				Val: 7,
			},
		},
	}

	fmt.Println("pre")
	binarytree.PrintTraversal(binarytree.TraversePreorder(&tree))
	fmt.Println("in")
	binarytree.PrintTraversal(binarytree.TraverseInorder(&tree))
	fmt.Println("post")
	binarytree.PrintTraversal(binarytree.TraversePostorder(&tree))
}
