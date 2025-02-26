package skiena_4

import (
	"github.com/greymatter-io/golangz/arrays"
	"testing"
)

// Devise an algorithm for finding the k smallest elements of an unsorted et of n integers in
// O(n + k(log n))

//Naive answer. Sort and choose the elements below the kth element from the array
//but cost is O(n log n) instead of O(n + k(log n).

//Right answer:
// Use your heap and iterate through the array of ints, pushing each element on to a heap with the
// heap being a min heap(a single function governs whether its a min or max heap).  Keep track
// of the size of the heap with a customization to your heap implementation.
// Every time the size of the heap exceeds k, delete the max element from the heap.
// You will end up with the kth smallest elements in the heap.

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
	for i := 0; i < len(h.hp)-1; i++ {
		x, _ := Pop(h)
		actual[i] = x
	}
	expected := []int{2, -3, 1}
	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Arrays not equal")
	}

}
