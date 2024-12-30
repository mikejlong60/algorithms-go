package chapter4

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type Frequency struct {
	probability float32
	letter      string
	l           *Frequency
	r           *Frequency
}

func (w Frequency) String() string {
	return fmt.Sprintf("Frequency{probability:%v, letter:%v, l:%v, r:%v}", w.probability, w.letter, w.l, w.r)
}

// Gets and deletes first element. Modifies the passed heap by deleteing the first element and restoring the heap property
// Returns - The heap with the top element deleted, the popped top element, and an error if the passed heap is empty
func HeapPopF(heap []*Frequency, lt func(l, r *Frequency) bool) ([]*Frequency, *Frequency, error) {
	fst, err := FindMinF(heap)
	if err != nil {
		return heap, nil, err
	} else {
		heap, err = HeapDeleteF(heap, 0, lt)
		if err != nil {
			return heap, nil, err
		}
		return heap, fst, nil
	}
}

// freqHeap and encodingHeap are the same starting out
// Returns - An error if the passed heap element index is greater than the weight of the heap
func Huffman(freqHeap []*Frequency, lt func(l, r *Frequency) bool) []*Frequency {
	if len(freqHeap) == 2 {
		freqHeap, fst, err := HeapPopF(freqHeap, lt)
		if err != nil {
			log.Errorf("HeapPopF failed:%v", err)
		}
		freqHeap, snd, err := HeapPopF(freqHeap, lt)
		if err != nil {
			log.Errorf("HeapPopF failed:%v", err)
		}

		meta := Frequency{
			probability: fst.probability + snd.probability,
			letter:      fmt.Sprintf("(%v:%v)", fst.letter, snd.letter),
			l:           fst,
			r:           snd,
		}
		return HeapInsertF(freqHeap, &meta, lt)
	} else {
		freqHeap, fst, err := HeapPopF(freqHeap, lt)
		if err != nil {
			log.Errorf("HeapPopF failed:%v", err)
		}
		freqHeap, snd, err := HeapPopF(freqHeap, lt)
		if err != nil {
			log.Errorf("HeapPopF failed:%v", err)
		}
		meta := Frequency{
			probability: fst.probability + snd.probability,
			letter:      fmt.Sprintf("(%v:%v)", fst.letter, snd.letter),
			l:           fst,
			r:           snd,
		}
		freqHeap = HeapInsertF(freqHeap, &meta, lt)
		return Huffman(freqHeap, lt)
	}
	//return freqHeap
}

// i int - the index in the given heap of the parent of element i. Array indices start with the number zero.
// Performance - O(1)
func FrequencyIdx(i int) int {
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
func HeapifyUpF(heap []*Frequency, i int, lt func(l, r *Frequency) bool) []*Frequency {
	if len(heap) == 0 {
		return heap
	}
	if i > 0 {
		j := FrequencyIdx(i)
		if lt(heap[i], heap[j]) {
			//Swap elements
			temp := heap[i]
			temp2 := heap[j]
			heap[j] = temp
			heap[i] = temp2
			heap = HeapifyUpF(heap, j, lt)
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
func HeapifyDownF(heap []*Frequency, i int, lt func(l, r *Frequency) bool) []*Frequency {
	var j int
	var n = len(heap) - 1
	if 2*i > n {
		return heap

	} else if 2*i < n {
		j = 0
		left := 2 * i
		var right = (2 * i) + 1
		leftVal := heap[left]
		rightVal := heap[right]
		if lt(leftVal, rightVal) {
			j = left
		} else {
			j = right
		}
	} else if 2*i == n {
		j = 2 * i
	}
	if lt(heap[j], heap[i]) {
		//Swap elements
		temp := heap[i]
		temp2 := heap[j]
		heap[j] = temp
		heap[i] = temp2
		heap = HeapifyDownF(heap, j, lt)
	}
	return heap
}

// Parameters:
//
//	n int - the size of the heap. This is fixed.
//
// Returns - A new heap (as a slice) of size n that has every element initialized to the zero value
// Performance - O(N)
func StartHeapF(n int) []*Frequency {
	var x = make([]*Frequency, n)
	return x
}

// This is a pure function.
// Parameters:
//
//	heap []A - the slice that is holding the heap
//
// Returns -the minimum element in the given heap without removing it. O(1)
// Performance - O(1)
func FindMinF(heap []*Frequency) (*Frequency, error) {
	if len(heap) == 0 || heap[0] == nil {
		return nil, fmt.Errorf("heap is empty or filled with nil pointers. FindMin is therefore irrelevant.")
	}
	return heap[0], nil
}

//	h []A - the slice that is holding the heap
//
// Returns - A pure function. The index of the first empty slot in the heap, or -1 if there are no empty slots
// Performance - O(N)
// TODO maybe you never have empty slots because you are using slices
func findFirstEmptySlotInHeapF(h []*Frequency) int {
	for i, x := range h {
		if x == nil {
			return i
		}
	}
	return -1
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
func HeapInsertF(heap []*Frequency, a *Frequency, lt func(l, r *Frequency) bool) []*Frequency {
	if len(heap) == 0 {
		heap = make([]*Frequency, 1)
	}

	l := findFirstEmptySlotInHeapF(heap)
	if l == -1 { //No empty slot so add a new spot at the end and put the new element there
		heap = append(heap, nil)
		l = len(heap) - 1
	}
	heap[l] = a
	return HeapifyUpF(heap, l, lt)
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
func HeapDeleteF(heap []*Frequency, i int, lt func(l, r *Frequency) bool) ([]*Frequency, error) {

	n := len(heap)
	if n == 0 {
		return []*Frequency{}, nil
	}

	if i > len(heap)-1 {
		log.Errorf("The element:%v you are trying to delete is longer than heap weight: %v", i, len(heap)-1)
		return heap, fmt.Errorf("The element:%v you are trying to delete is longer than heap weight: %v", i, len(heap)-1)
	}

	//Delete last and only element from heap
	if i == len(heap)-1 {
		return []*Frequency{}, nil
	}
	heap[i] = heap[len(heap)-1]
	heap = heap[0 : len(heap)-1]
	if len(heap) == 1 {
		return heap, nil
	} else {
		return HeapifyDownF(heap, i, lt), nil
	}
}
