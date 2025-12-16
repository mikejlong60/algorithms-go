package skiena_4

import (
	"fmt"
	"testing"
	"time"

	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/heap"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/mikejlong60/algorithms-go/chapter5"
	log "github.com/sirupsen/logrus"
)

// Show that you can sort an array of k distinct integers in O(n log k) steps, better than O(n log n).  Think of sorting
// an array of 10,000 ones and zeros.

//This can be shown to be true using a heap. Since any value can be inserted into a heap in log n steps where n is the
//size of the heap.  In the case of the examole above where k is two, the algorithm is as follows:

/*
*
heap y
for i, _ := range xs { // O(log k) * n total steps.

		incrementI(y, i)
	}

Now just make a new array by popping off the heap until it is empty and writing number of increment elements for k.

I think this must be how heap sort is implemented.
*/
func lt(x, y *int) bool {
	if *x < *y {
		return true
	} else {
		return false
	}
}

var elementBExtractor = func(c *int) int {
	return *c
}

func heapSortInsertIntoHeap(xss []int) propcheck.Pair[heap.Heap[int, int], map[int]int] {
	elementCount := make(map[int]int)
	var h = heap.New[int, int](elementBExtractor)
	for _, x := range xss {
		p := heap.FindPosition(h, x)
		if p < 0 { //element is not in heap yet
			elementCount[x] = 1
			h = heap.HeapInsert(h, &x, lt)
		} else { //increment count of element's appearance in array
			elementCount[x] = elementCount[x] + 1
		}
	}
	return propcheck.Pair[heap.Heap[int, int], map[int]int]{h, elementCount}
}

func TestHeapDeleteEveryElementStartingFromLast(t *testing.T) {
	var sortInOLogK = func(xss []int) propcheck.Pair[[]int, []int] {
		result := make([]int, 0)
		start := time.Now()

		p := heapSortInsertIntoHeap(xss)
		var h = p.A
		var err error
		keyCount := p.B
		for {
			m, _ := heap.FindMin(p.A)
			h, err = heap.HeapDelete(h, 0, lt)
			if err != nil {
				break
			}
			count := keyCount[*m]
			iresult := make([]int, count)
			for i := 0; i < count; i++ {
				iresult[i] = *m
			}
			result = append(result, iresult...)
		}
		log.Infof("heapsort array of:%v ints took:%v", len(xss), time.Since(start))
		return propcheck.Pair[[]int, []int]{result, xss}
	}

	compareSpeedWithMergeSort := func(p propcheck.Pair[[]int, []int]) (bool, error) {

		lt := func(x, y int) bool {
			if x < y {
				return true
			} else {
				return false
			}
		}

		expected := make([]int, len(p.B))
		copy(expected, p.B)
		start := time.Now()
		expected = chapter5.MergeSort(expected, lt)
		log.Infof("my mergesort array of:%v ints took:%v", len(p.B), time.Since(start))

		var errors error
		if !arrays.ArrayEquality(p.A, expected, eq) {
			errors = fmt.Errorf("Actual:%v Expected:%v", p.A, expected)
		}

		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}

	ge := propcheck.ChooseInt(0, 3)
	g0 := propcheck.ChooseArray(100000, 100000, ge)
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	prop := propcheck.ForAll(g0,
		"Validate HeapDelete  \n",
		sortInOLogK,
		compareSpeedWithMergeSort,
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]int](t, result)
}
