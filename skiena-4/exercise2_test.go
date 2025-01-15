package skiena_4

import (
	"fmt"
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/mikejlong60/algorithms-go/chapter5"
	"math"
	"testing"
)

func TestMaximumSpaceBetweenTwoElementsSorted(t *testing.T) {
	lt := func(l, r int) bool {
		if l < r {
			return true
		} else {
			return false
		}
	}

	r := []int{1, 1, 13, 14, 234, 43, 13, 45, 15, 26, 7, 18, 90, 1, 12, 23}
	rr := chapter5.MergeSort(r, lt)
	//Big O(1) since list was sorted previously
	actualMaxDistance := rr[len(rr)-1] - rr[0]
	expectedMaxDistance := 233
	if actualMaxDistance != expectedMaxDistance {
		t.Errorf("\nActual:  %v\nExpected:%v", actualMaxDistance, expectedMaxDistance)
	}
}

func TestMaximumSpaceBetweenTwoElementsUnsorted(t *testing.T) {
	f := func(minMax propcheck.Pair[int, int], x int) propcheck.Pair[int, int] {
		if x < minMax.A {
			minMax.A = x
		}
		if x > minMax.B {
			minMax.B = x
		}
		return minMax
	}
	//Big O(n)
	r := []int{1, 1, 13, 14, 234, 43, 13, 45, 15, 26, 7, 18, 90, 1, 12, 23}
	rr := arrays.FoldLeft(r, propcheck.Pair[int, int]{math.MaxInt, math.MinInt}, f)
	actualMaxDistance := rr.B - rr.A
	expectedMaxDistance := 233
	if actualMaxDistance != expectedMaxDistance {
		t.Errorf("\nActual:  %v\nExpected:%v", actualMaxDistance, expectedMaxDistance)
	}
}

func TestMinimumSpaceBetweenTwoElementsSorted(t *testing.T) {
	rr := []int{1, 3, 4, 5, 5, 6}

	actualMinDistancePair := propcheck.Pair[int, int]{0, math.MaxInt}
	fmt.Println(actualMinDistancePair.B - actualMinDistancePair.A)
	for i := 0; i < len(rr)-1; i++ {
		if rr[i+1]-rr[i] < actualMinDistancePair.B-actualMinDistancePair.A {
			actualMinDistancePair = propcheck.Pair[int, int]{rr[i], rr[i+1]}
		}
	}

	expectedMinDistancePair := propcheck.Pair[int, int]{5, 5}
	if !(actualMinDistancePair.A == expectedMinDistancePair.A && actualMinDistancePair.B == expectedMinDistancePair.B) {
		t.Errorf("\nActual:  %v\nExpected:%v", actualMinDistancePair, expectedMinDistancePair)
	}
}

func TestMinimumSpaceBetweenTwoElementsUnSorted(t *testing.T) {
	lt := func(l, r int) bool {
		if l < r {
			return true
		} else {
			return false
		}
	}

	//Big O(n)
	r := []int{1, 3, 5, 4, 5, 6}
	rr := chapter5.MergeSort(r, lt)
	actualMinDistancePair := propcheck.Pair[int, int]{0, math.MaxInt}
	fmt.Println(actualMinDistancePair.B - actualMinDistancePair.A)
	for i := 0; i < len(rr)-1; i++ {
		if rr[i+1]-rr[i] < actualMinDistancePair.B-actualMinDistancePair.A {
			actualMinDistancePair = propcheck.Pair[int, int]{rr[i], rr[i+1]}
		}
	}

	expectedMinDistancePair := propcheck.Pair[int, int]{5, 5}
	if !(actualMinDistancePair.A == expectedMinDistancePair.A && actualMinDistancePair.B == expectedMinDistancePair.B) {
		t.Errorf("\nActual:  %v\nExpected:%v", actualMinDistancePair, expectedMinDistancePair)
	}
}
