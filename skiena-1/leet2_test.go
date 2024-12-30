package skiena_1

import (
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/linked_list"
	"testing"
)

// Not a pure function
func rotateList(xs *linked_list.LinkedList[int], k, soFar int) *linked_list.LinkedList[int] {
	if k == soFar {
		return xs
	} else {
		x := linked_list.Head(xs)
		xs, _ := linked_list.Tail(xs)
		xs = linked_list.AddLast(x, xs)
		return rotateList(xs, k, soFar+1)
	}
}

func TestRotateList(t *testing.T) {
	eq := func(l, r int) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}
	expected := []int{3, 4, 5, 1, 2}
	xs := linked_list.ToList([]int{1, 2, 3, 4, 5})
	actual := rotateList(xs, 2, 0)
	if !arrays.ArrayEquality(linked_list.ToArray(actual), expected, eq) {
		t.Errorf("Actual:%v Expected:%v", xs, expected)
	}

}
