package skiena_4

import (
	"github.com/greymatter-io/golangz/arrays"
	"testing"
)

// Give an O(n log k) algorithm that merges k sorted lists with a total of n elements in one
// sorted list.

func insertIntoHeap2(xss []int, h Heap) Heap {
	for _, x := range xss {
		h = HeapInsert(h, x)
	}
	return h
}

func TestMergeKSortedArrays(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
	c := []int{26, 27, 28, 29, 30, 31, 32, 33}
	d := [][]int{a, b, c}
	h := New()
	for _, x := range d {
		h = insertIntoHeap2(x, h)
	}
	var actual = make([]int, 33)
	var x int
	for i := 0; i < len(h.hp)-1; i++ {
		x, h, _ = Pop(h)
		actual[i] = x
	}
	expected := []int{2, -3, 1}
	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Arrays not equal")
	}

}
