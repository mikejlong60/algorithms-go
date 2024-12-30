package skiena_3

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
	"github.com/hashicorp/go-multierror"
	"math"
	"testing"
	"time"
)

var currentDepth int
var maxDepth int

func traverse[T any](btree *Node[T]) int {

	if btree == nil {
		oldMaxDepth := currentDepth
		currentDepth = 0
		if oldMaxDepth > maxDepth {
			maxDepth = oldMaxDepth
		}
		return maxDepth
	} else {
		currentDepth = currentDepth + 1
		traverse(btree.left)
		traverse(btree.right)
	}
	return maxDepth
}

func TestMaxDepth(t *testing.T) {

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
	g0 := sets.ChooseSet(1, 60000, propcheck.ChooseInt(-1000000, 1000000), lt, eq)

	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	prop := propcheck.ForAll(g0,
		"Find Max depth of binary tree. Tree does not have to be height-balanced",
		func(a []int) []int {
			currentDepth = 0
			maxDepth = 0
			return a
		},
		func(a []int) (bool, error) {

			btree := BinaryTree(a, lt)
			var errors error

			actual := traverse(btree)
			expected := int(math.Log2(float64(len(a))))

			if !(actual-1 <= expected && actual+1 >= expected) {
				errors = multierror.Append(errors, fmt.Errorf("Actual:%v Expected:%v", actual, expected))
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
