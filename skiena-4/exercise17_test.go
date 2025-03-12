package skiena_4

import (
	"fmt"
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sorting"
	"github.com/hashicorp/go-multierror"
	log "github.com/sirupsen/logrus"
	"math"
	"testing"
	"time"
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

type Heap struct {
	hp []int
}

func New() Heap {
	return Heap{
		hp: make([]int, 0),
	}
}

// i int - the index in the given heap of the parent of element i. Array indices start with the number zero.
// Performance - O(1)
func ParentIdx(i int) int {
	//Odd number
	if i%2 > 0 {
		return i / 2
	} else { // even number
		return (i / 2) - 1
	}
}

// Definition of almost-a-heap. Only one node in the tree has a value less than it's parent as per the gt function and that
// node is at the bottom rung of the heap.
// Definition of a heap.  Every node in the tree has a greater value than it's parent as per the gt function.
// This is a not pure function
// Parameters:
//
//	h - the heap object containing the heap(represented as a slice).
//	i int - the index into the heap of the element you want to move up. Array indices start with the number zero.
//	gt func(l, r A) bool - A predicate function that determines whether the left A element is less than the right A element.
//
// Returns - The modified heap that has the i'th element in its proper position in the heap
// Performance - O(log N) assuming that the array is almost-a-heap with the key: heap(i) too small.
func heapifyUp(h Heap, i int) Heap {
	if len(h.hp) == 0 {
		return h
	}
	if i > 0 {
		j := ParentIdx(i)
		if gt(h.hp[i], h.hp[j]) {
			//Swap elements
			temp := h.hp[i]
			temp2 := h.hp[j]

			h.hp[j] = temp
			h.hp[i] = temp2
			h = heapifyUp(h, j)
		}
	}
	return h
}

// This is not a pure function because it modifies the array each time.
// Parameters:
//
//	h - the heap object containing the heap(represented as a slice).
//	i int - the index into the heap of the element you want to move up. Array indices start with the number zero.
//	gt func(l, r A) bool - A predicate function that determines whether the left A element is less than the right A element.
//
// Returns - The modified heap that has the i'th element in its proper position in the heap
// Performance - O(log N) assuming that the array is almost-a-heap with the key: heap(i) too big.
func heapifyDown(h Heap, i int) Heap {
	var j int
	n := len(h.hp)
	if (2*i)+1 > n {
		return h
	} else if (2*i)+1 < n {
		j = 0
		//These differ from book definition because array indices start with zero
		left := (2 * i) + 1
		right := (2 * i) + 2
		leftVal := h.hp[left]
		if right < n {
			rightVal := h.hp[right]
			if gt(leftVal, rightVal) {
				j = left
			} else {
				j = right
			}
		} else {
			j = left
		}
	} else if (2*i)+1 == n {
		j = (2 * i) + 1
	}
	if j < n && gt(h.hp[j], h.hp[i]) {
		//Swap elements
		temp := h.hp[i]
		temp2 := h.hp[j]
		h.hp[j] = temp
		h.hp[i] = temp2
		h = heapifyDown(h, j)
	}
	return h
}

// This is a pure function.
// Parameters:
//
//	h - the heap object containing the heap(represented as a slice).
//
// Returns -the minimum element in the given heap without removing it. O(1)
// Performance - O(1)
// This is a hack but 0 is the default value for ints.
// TODO make it min int instead.
func FindMin(h Heap) (int, error) {
	if len(h.hp) == 0 {
		return -1, fmt.Errorf("heap is empty so findMin is therefore irrelevant")
	}
	return h.hp[0], nil
}

func Pop(h Heap) (int, Heap, error) {
	x, er := FindMin(h)
	if er != nil {
		return math.MinInt, h, er
	}
	h, er = HeapDelete(h, 0)
	if er != nil {
		return math.MinInt, h, er
	}
	return x, h, nil
}

// Inserts the given element into the given heap and returns the modified heap.
//
// O(log n)
//
// Parameters:
//
//	h - the heap object containing the heap(represented as a slice).
//	a  A - the element you want to insert into the heap
//	gt func(l, r A) bool - A predicate function that determines whether or not the left A element is less than the right A element.
//
// Returns - The original heap (as a slice) that has the given element in its proper position
// Performance - O(log N)
// NOTE This function assumes that the heap slice has no empty elements. It always adds a new one.
func HeapInsert(h Heap, a int) Heap {
	h.hp = append(h.hp, 0) //Adds an empty element at end
	l := len(h.hp) - 1     //Get index of end of heap and stick new element there
	h.hp[l] = a

	//Now move it up as necessary until that part of tree satisfies heap property
	return heapifyUp(h, l)
}

// Determines if given heap is empty.
//
//	h - the heap object containing the heap(represented as a slice).
//
// Returns - Whether or not the heap is empty
// Performance - O(1)
func Empty(h Heap) bool {
	if len(h.hp) == 0 {
		return true
	} else {
		return false
	}
}

// Deletes an element from the given heap. This is not a pure function.
// Parameters:
//
//	h - the heap object containing the heap(represented as a slice).
//	i int - the index into the heap of the element you want to delete. Array indices start with the number zero.
//	gt func(l, r A) bool - A predicate function that determines whether or not the left A element is less than the right A element.
//
// Returns - The original heap that has the given element in its proper position.
//
//	If the heao is empty or the indice you are trying to delete is longer than the heap(zero indexed) then you get an error
//
// Performance - O(log N)
func HeapDelete(h Heap, i int) (Heap, error) {
	if i > len(h.hp)-1 || len(h.hp) == 0 {
		log.Errorf("element:%v you are trying to delete is longer than heap length: %v", i, len(h.hp)-1)
		return h, fmt.Errorf("element:%v you are trying to delete is longer than heap length: %v", i, len(h.hp)-1)
	}

	//Delete last element and return. No need to move anything around.
	if i == len(h.hp)-1 {
		h.hp = h.hp[0 : len(h.hp)-1]
		return h, nil
	}
	h.hp[i] = h.hp[len(h.hp)-1]
	h.hp = h.hp[0 : len(h.hp)-1]

	if len(h.hp) == 1 {
		return h, nil
	}

	parent := ParentIdx(i)
	if parent > 0 && gt(h.hp[i], h.hp[parent]) {
		return heapifyUp(h, i), nil
	} else {
		return heapifyDown(h, i), nil
	}
}

func gt(l, r int) bool {
	if l > r {
		return true
	} else {
		return false
	}
}

func eq(l, r int) bool {
	if l == r {
		return true
	} else {
		return false
	}
}

func insertIntoHeap(xss []int) Heap {
	var h = New()
	for _, x := range xss {
		h = HeapInsert(h, x)
	}
	return h
}
func parentIsLessThanOrEqual(h Heap, lastIdx int) error {
	var pIdx = ParentIdx(lastIdx)
	var cIdx = lastIdx
	var errors error
	for pIdx >= 0 {
		if !gt(h.hp[pIdx], h.hp[cIdx]) {
			errors = multierror.Append(errors, fmt.Errorf("parent:%v value was not less than or equal to child's:%v\n", h.hp[pIdx], h.hp[cIdx]))
		}
		cIdx = pIdx
		pIdx = ParentIdx(cIdx)
	}
	return errors
}

func validateIsAHeap(p Heap) (bool, error) {
	var errors error
	for idx, _ := range p.hp {
		errors = parentIsLessThanOrEqual(p, idx)
	}
	if errors != nil {
		return false, errors
	} else {
		return true, nil
	}
}

// Find the minimum and compares it to the actual min in the initial array.
// If array is empty or filled with nil pointers that is OK and FindMin should not fail
// but return a Golang error, not panic.
func minimumCorrectValue(h Heap, sorted []int) bool {
	key, err := FindMin(h)
	if len(h.hp) > 0 && err == nil {
		return eq(key, sorted[0])
	} else {
		return true
	}
}
func validateHeapMin(p Heap) (bool, error) {
	var errors error
	var sorted = make([]int, len(p.hp))
	copy(sorted, p.hp)
	sorting.QuickSort(sorted, gt)
	if !minimumCorrectValue(p, sorted) {
		errors = multierror.Append(errors, fmt.Errorf("FindMin should have returned:%v", sorted[0]))
	}
	if errors != nil {
		return false, errors
	} else {
		return true, nil
	}
}

func TestHeapInsertWithEmptyHeap(t *testing.T) {
	g := propcheck.ChooseArray(0, 1000, propcheck.ChooseInt(0, 10000000))
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}

	prop := propcheck.ForAll(g,
		"Validate heapifyUp  \n",
		insertIntoHeap,
		validateIsAHeap, validateHeapMin,
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]int](t, result)
}

func TestFindKSmallestElements(t *testing.T) {

	findKSmallestElements := func(xss []int, k int) Heap {
		heapSize := 0
		var h = New()
		for _, x := range xss {
			if heapSize > k {
				h, _ = HeapDelete(h, 0)
				heapSize--
			}
			h = HeapInsert(h, x)
			heapSize++
		}
		h, _ = HeapDelete(h, 0) //delete last extra element
		return h
	}

	a := []int{3, 2, 1, 13, 5, 6, 7, 9, 11, -3}
	actual := findKSmallestElements(a, 3)
	expected := []int{2, -3, 1}
	if !arrays.ArrayEquality(actual.hp, expected, eq) {
		t.Errorf("Arrays not equal")
	}

	actual = findKSmallestElements(a, 4)
	expected = []int{3, 2, 1, -3}
	if !arrays.ArrayEquality(actual.hp, expected, eq) {
		t.Errorf("Arrays not equal")
	}
	a = []int{3, 2, 1, 21, 12, 8, 13, 5, 6, 7, 9, 11}
	actual = findKSmallestElements(a, 4)
	expected = []int{5, 3, 1, 2}
	if !arrays.ArrayEquality(actual.hp, expected, eq) {
		t.Errorf("Arrays not equal")
	}

}
