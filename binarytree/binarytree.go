package binarytree

import "fmt"

type Node[T int | string] struct {
	Val   T
	Left  *Node[T]
	Right *Node[T]
}

func TraversePreorder[T int | string](root *Node[T]) []T {
	if root == nil {
		return nil
	}

	left := TraversePreorder(root.Left)
	right := TraversePreorder(root.Right)

	return append(append([]T{root.Val}, left...), right...)
}

func TraverseInorder[T int | string](root *Node[T]) []T {
	if root == nil {
		return nil
	}

	left := TraverseInorder(root.Left)
	right := TraverseInorder(root.Right)

	return append(append(left, root.Val), right...)
}

func TraversePostorder[T int | string](root *Node[T]) []T {
	if root == nil {
		return nil
	}

	left := TraversePostorder(root.Left)
	right := TraversePostorder(root.Right)

	return append(append(left, right...), root.Val)
}

func PrintTraversal[T int | string](nodeVals []T) {
	for i := range nodeVals {
		fmt.Println(nodeVals[i])
	}
}
