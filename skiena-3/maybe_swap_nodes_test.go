package skiena_3

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
	"github.com/hashicorp/go-multierror"
	"testing"
	"time"
)

func swap[T any](btree *Node[T]) {
	tleft := btree.left
	btree.left = btree.right
	btree.right = tleft
}
func maybeSwap[T any](btree *Node[T], numSwapped int, lt func(l, r T) bool) int {

	if btree == nil {
		return numSwapped
	} else {
		if btree.left != nil && btree.right != nil {
			if !lt(btree.left.val, btree.right.val) {
				swap(btree)
				numSwapped = numSwapped + 1
			}
		}
		numSwapped = maybeSwap(btree.left, numSwapped, lt)
		numSwapped = maybeSwap(btree.right, numSwapped, lt)
	}
	return numSwapped
}

func TestMaybeSwap(t *testing.T) {

	lt := func(l, r int) bool {
		if l < r {
			return true
		} else {
			return false
		}
	}

	eq := func(l, r int) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}
	g0 := sets.ChooseSet(7, 100, propcheck.ChooseInt(-1000000, 1000000), lt, eq)

	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	prop := propcheck.ForAll(g0,
		"Swap two nodes of binary tree if they do not meet the predicate test. Tree does not have to be height-balanced",
		func(a []int) *Node[int] {
			btree := BinaryTree(a, lt)
			swap(btree.left) //screw up the tree to test
			return btree
		},
		func(btree *Node[int]) (bool, error) {

			var errors error

			actual := maybeSwap(btree, 0, lt)
			if actual != 1 {
				errors = multierror.Append(errors, fmt.Errorf("Actual:%v Expected:%v", actual, 1))
			}
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]int](t, result)
}
