package chapter2

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type Cache struct {
	ts   int
	data string
}

func lt(l, r *Cache) bool {
	if l != nil && r != nil && l.ts < r.ts {
		return true
	} else {
		return false
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

// Definition of almost-a-heap. Only one node in the tree has a value less than it's parent as per the lt function and that
// node is at the bottom rung of the heap.
// Definition of a heap.  Every node in the tree has a greater value than it's parent as per the lt function.
// This is a not pure function
// Parameters:
//
//	heap []A - the slice that is holding the heap
//	i int - the index into the heap of the element you want to move up. Array indices start with the number zero.
//	lt func(l, r A) bool - A predicate function that determines whether or not the left A element is less than the right A element.
//
// Returns - The modified heap (as a slice) that has the i'th element in its proper position in the heap
// Performance - O(log N) assuming that the array is almost-a-heap with the key: heap(i) too small.
func HeapifyUp(heap []*Cache, i int, lt func(l, r *Cache) bool) []*Cache {
	if len(heap) == 0 {
		return []*Cache{}
	}
	if i > 0 {
		j := ParentIdx(i)
		if lt(heap[i], heap[j]) {
			//Swap elements
			temp := heap[i]
			temp2 := heap[j]
			heap[j] = temp
			heap[i] = temp2
			heap = HeapifyUp(heap, j, lt)
		}
	}
	return heap
}

// This is not a pure function because it modified the array each time.
// Parameters:
//
//	heap []A - the slice that is holding the heap
//	i int - the index into the heap of the element you want to move down. Array indices start with the number zero.TODO change
//	lt func(l, r A) bool - A predicate function that determines whether or not the left A element is less than the right A element.
//
// Returns - The original heap (as a slice) that has the i'th element in its proper position
// Performance - O(log N) assuming that the array is almost-a-heap with the key: heap(i) too big.
func HeapifyDown(heap []*Cache, i int, lt func(l, r *Cache) bool) []*Cache {
	var j int
	n := len(heap)
	if (2*i)+1 > n {
		return heap

	} else if (2*i)+1 < n {
		j = 0
		//These differ from book definition because array indices start with zero
		left := (2 * i) + 1
		right := (2 * i) + 2
		leftVal := heap[left]
		if right < n {
			rightVal := heap[right]
			if lt(leftVal, rightVal) {
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
	if j < n && lt(heap[j], heap[i]) {
		//Swap elements
		temp := heap[i]
		temp2 := heap[j]
		heap[j] = temp
		heap[i] = temp2
		heap = HeapifyDown(heap, j, lt)
	}
	return heap
}

// This is a pure function.
// Parameters:
//
//	heap []A - the slice that is holding the heap
//
// Returns -the minimum element in the given heap without removing it. O(1)
// Performance - O(1)
func FindMin(heap []*Cache) (*Cache, error) {
	if len(heap) == 0 || heap[0] == nil {
		return nil, fmt.Errorf("heap is empty. FindMin is therefore irrelevant.")
	}
	return heap[0], nil
}

// Inserts the given element into the given heap and returns the modified heap.
//
// O(log n)
//
// Parameters:
//
//	heap []A - the slice that is holding the heap
//	a  A - the element you want to insert into the heap
//	lt func(l, r A) bool - A predicate function that determines whether or not the left A element is less than the right A element.
//
// Returns - The original heap (as a slice) that has the given element in its proper position
// Performance - O(log N)
// NOTE This function assumes that the heap slice has no empty elements. It always adds a new one.
func HeapInsert(heap []*Cache, a *Cache, lt func(l, r *Cache) bool) []*Cache {
	heap = append(heap, nil)
	l := len(heap) - 1
	heap[l] = a
	return HeapifyUp(heap, l, lt)
}

// Deletes an element from the given heap. This is not a pure function.
// Parameters:
//
//	heap []A - the slice that is holding the heap
//	i int - the index into the heap of the element you want to delete. Array indices start with the number zero.
//	lt func(l, r A) bool - A predicate function that determines whether or not the left A element is less than the right A element.
//
// Returns - The original heap that has the given element in its proper position
// Performance - O(log N)
func HeapDelete(heap []*Cache, i int, lt func(l, r *Cache) bool) ([]*Cache, error) {

	n := len(heap)
	if n == 0 {
		return []*Cache{}, nil
	}

	if i > len(heap)-1 {
		log.Errorf("The element:%v you are trying to delete is longer than heap length: %v", i, len(heap)-1)
		return heap, fmt.Errorf("The element:%v you are trying to delete is longer than heap length: %v", i, len(heap)-1)
	}

	//Delete last and only element from heap
	if i == len(heap)-1 {
		return []*Cache{}, nil
	}
	heap[i] = heap[len(heap)-1]
	heap = heap[0 : len(heap)-1]
	if len(heap) == 1 {
		return heap, nil
	}

	parent := ParentIdx(i)
	if parent > 0 && lt(heap[i], heap[parent]) {
		return HeapifyUp(heap, i, lt), nil
	} else {
		return HeapifyDown(heap, i, lt), nil
	}
}
