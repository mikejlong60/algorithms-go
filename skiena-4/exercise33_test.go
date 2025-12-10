package skiena_4

import (
	"fmt"
	"testing"
	"time"

	"github.com/greymatter-io/golangz/heap"
	"github.com/greymatter-io/golangz/propcheck"
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

type HowMany struct {
	key   int
	count int
}

func lt(x, y *HowMany) bool {
	if x.key < y.key {
		return true
	} else {
		return false
	}
}

var elementBExtractor = func(c *HowMany) int {
	return c.key
}

func heapSortInsertIntoHeap(xss []int) heap.Heap[HowMany, int] {
	var h = heap.New[HowMany](elementBExtractor)
	for _, x := range xss {
		h = heap.HeapInsert(h, &HowMany{x, 0}, lt)
	}
	return h
}

func TestHeapDeleteEveryElementStartingFromLast(t *testing.T) {
	var sortInOLogK = func(xss []int) []HowMany {
		result := make([]HowMany, 0)
		var h = heapSortInsertIntoHeap(xss)
		for _ = range xss {
			m, err := heap.FindMin(h)
			if err != nil {
				fmt.Println(err)
				break
			}
			m.count = m.count + 1
			result = append(result, *m)
			h, _ = heap.HeapDelete(h, 0, lt)
		}
		return result
	}
	trueId := func(p []HowMany) (bool, error) {
		return true, nil
	}

	ge := propcheck.ChooseInt(0, 3)
	g0 := propcheck.ChooseArray(100, 500, ge)
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	prop := propcheck.ForAll(g0,
		"Validate HeapDelete  \n",
		sortInOLogK,
		trueId,
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]int](t, result)
}
