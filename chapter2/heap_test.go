package chapter2

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sorting"
	"github.com/hashicorp/go-multierror"
	"testing"
	"time"
)

func gtWhenNonDefaultChild(l, r *Cache) bool {
	if l.ts > r.ts && r != nil {
		return true
	} else {
		return false
	}
}

// Find the mimimum and compares it to the actual min in the initial array.
// If array is empty or filled with nil pointers that is OK and FindMin should not fail
// but return a Golang error, not panic.
func minimumCorrectValue(p, sorted []*Cache) bool {
	min, err := FindMin(p)
	if len(p) > 0 && err == nil {
		return min.ts == sorted[0].ts
	} else {
		return true
	}
}

func parentIsLessThanOrEqual(heap []*Cache, lastIdx int, parentGT func(l, r *Cache) bool) error {
	var pIdx = ParentIdx(lastIdx)
	var cIdx = lastIdx
	var errors error
	for pIdx > 0 {
		if parentGT(heap[pIdx], heap[cIdx]) {
			errors = multierror.Append(errors, fmt.Errorf("parent:%v value was not less than or equal to child's:%v\n", heap[pIdx], heap[cIdx]))
		}
		cIdx = pIdx
		pIdx = ParentIdx(cIdx)
	}
	return errors
}

func TestHeapInsertWithEmptyHeap(t *testing.T) {
	g := propcheck.ChooseArray(0, 5, propcheck.ChooseInt(0, 10000))
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	insertIntoHeap := func(xss []int) []*Cache {
		var r = []*Cache{}
		for _, x := range xss {
			r = HeapInsert(r, &Cache{x, fmt.Sprintf("ts:%v", x)}, lt)
		}
		return r
	}
	insert := func(p []int) []*Cache {
		xss := insertIntoHeap(p)
		return xss
	}
	validateIsAHeap := func(p []*Cache) (bool, error) {
		var errors error
		for idx := range p {
			errors = parentIsLessThanOrEqual(p, idx, gtWhenNonDefaultChild)
		}
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}
	validateHeapMin := func(p []*Cache) (bool, error) {
		var errors error
		var sorted = make([]*Cache, len(p))
		copy(sorted, p)
		sorting.QuickSort(sorted, lt)
		if !minimumCorrectValue(p, sorted) {
			errors = multierror.Append(errors, fmt.Errorf("FindMin should have returned:%v", sorted[0]))
		}
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}

	prop := propcheck.ForAll(g,
		"Validate HeapifyUp  \n",
		insert,
		validateIsAHeap, validateHeapMin,
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]int](t, result)
}

func TestHeapInsertWithNonEmptyHeapHeap(t *testing.T) {
	g := propcheck.ChooseArray(10, 1000, propcheck.ChooseInt(0, 10000))
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	insertIntoHeap := func(xss []int) []*Cache {
		var r = []*Cache{}
		for _, x := range xss {
			r = HeapInsert(r, &Cache{x, fmt.Sprintf("ts:%v", x)}, lt)
		}
		return r
	}
	insert := func(p []int) []*Cache {
		xss := insertIntoHeap(p)
		return xss
	}
	validateIsAHeap := func(p []*Cache) (bool, error) {
		var errors error
		for idx, x := range p {
			if x != nil { //Heap could have some empty elements at end
				errors = parentIsLessThanOrEqual(p, idx, gtWhenNonDefaultChild)
			}
		}
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}

	ltNoNilCheck := func(l, r *Cache) bool {
		if l.ts < r.ts {
			return true
		} else {
			return false
		}
	}

	validateHeapMin := func(p []*Cache) (bool, error) {
		var errors error
		var sorted = make([]*Cache, len(p))
		copy(sorted, p)
		for idx, x := range sorted { //Trim off em,pty elements at end to work around your sorting bug
			if x == nil { //Heap could have some empty elements at end
				sorted = sorted[0:idx]
				break
			}
		}

		sorting.QuickSort(sorted, ltNoNilCheck)
		if !minimumCorrectValue(p, sorted) {
			errors = multierror.Append(errors, fmt.Errorf("FindMin should have returned:%v", sorted[0]))
		}
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}

	prop := propcheck.ForAll(g,
		"Validate HeapifyUp  \n",
		insert,
		validateIsAHeap, validateHeapMin,
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]int](t, result)
}

func TestHeapDeleteSpecificElements(t *testing.T) {

	insertIntoHeap := func(xss []int) []*Cache {
		var r = []*Cache{}
		for _, x := range xss {
			r = HeapInsert(r, &Cache{x, fmt.Sprintf("ts:%v", x)}, lt)
		}
		return r
	}

	var delete6ElementsFromHeapOfAtLeast6 = func(xss []int) []*Cache {
		r := insertIntoHeap(xss)
		r, _ = HeapDelete(r, 5, lt)
		r, _ = HeapDelete(r, 4, lt)
		r, _ = HeapDelete(r, 3, lt)
		r, _ = HeapDelete(r, 2, lt)
		r, _ = HeapDelete(r, 1, lt)
		r, _ = HeapDelete(r, 0, lt)
		return r
	}

	validateIsAHeap := func(p []*Cache) (bool, error) {
		var errors error
		for idx := range p {
			errors = parentIsLessThanOrEqual(p, idx, gtWhenNonDefaultChild)
		}
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}
	validateHeapMin := func(p []*Cache) (bool, error) {
		var errors error
		var sorted = make([]*Cache, len(p))
		copy(sorted, p)
		sorting.QuickSort(sorted, lt)
		if !minimumCorrectValue(p, sorted) {
			errors = multierror.Append(errors, fmt.Errorf("FindMin should have returned:%v", sorted[0]))
		}
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}

	g0 := propcheck.ChooseArray(6, 15, propcheck.ChooseInt(1, 2000))
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	prop := propcheck.ForAll(g0,
		"Validate HeapDelete  \n",
		delete6ElementsFromHeapOfAtLeast6,
		validateIsAHeap, validateHeapMin,
	)
	result := prop.Run(propcheck.RunParms{1000, rng}) //The 3rd iteration paniced with array out or bounds
	propcheck.ExpectSuccess[[]int](t, result)
}

func TestHeapDeleteMinElement(t *testing.T) {
	var errors error
	insertIntoHeap := func(xss []int) []*Cache {
		var r = []*Cache{}
		for _, x := range xss {
			r = HeapInsert(r, &Cache{x, fmt.Sprintf("ts:%v", x)}, lt)
		}
		return r
	}

	correctHeapMin := func(p []*Cache) bool {
		var sorted = make([]*Cache, len(p))
		copy(sorted, p)
		sorting.QuickSort(sorted, lt)
		min, err := FindMin(p)
		if len(p) == 0 {
			return true
		} else if err != nil {
			return false
		} else if sorted[0].ts != min.ts {
			return false
		} else {
			return true
		}
	}

	var deleteAllFromHeap = func(xss []int) []*Cache {
		var r = insertIntoHeap(xss)
		for range r {
			r, _ = HeapDelete(r, 0, lt)
			if !correctHeapMin(r) {
				errors = multierror.Append(errors, fmt.Errorf("Heap property violated"))
			}
		}
		return r
	}

	validateIsAHeap := func(p []*Cache) (bool, error) {
		var errors error
		for idx := range p {
			errors = parentIsLessThanOrEqual(p, idx, gtWhenNonDefaultChild)
		}
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}

	heapWrong := func(p []*Cache) (bool, error) {
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}

	g0 := propcheck.ChooseArray(0, 1000, propcheck.ChooseInt(1, 200000))
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	prop := propcheck.ForAll(g0,
		"Validate HeapDelete  \n",
		deleteAllFromHeap,
		validateIsAHeap, heapWrong,
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]int](t, result)
}
