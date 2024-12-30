package skiena_3

import (
	"github.com/greymatter-io/golangz/option"
	"github.com/greymatter-io/golangz/sorting"
)

type Node[T any] struct {
	val   T
	left  *Node[T]
	right *Node[T]
}

type KeyPad struct {
	number  int
	letters []string
}

func Find[T any](btree *Node[T], val T, lt func(l, r T) bool, eq func(l, r T) bool) option.Option[T] {
	if btree == nil {
		return option.None[T]{}
	} else {
		if eq(btree.val, val) {
			return option.Some[T]{val}
		} else if lt(val, btree.val) {
			return Find(btree.left, val, lt, eq)
		} else {
			return Find(btree.right, val, lt, eq)
		}
	}
}

func BinaryTree[T any](arr []T, lt func(l, r T) bool) *Node[T] {

	sorting.QuickSort(arr, lt)

	if len(arr) == 0 {
		return nil
	}

	mid := len(arr) / 2
	root := &Node[T]{val: arr[mid]}
	root.left = BinaryTree[T](arr[:mid], lt)
	root.right = BinaryTree[T](arr[mid+1:], lt)

	return root
}
