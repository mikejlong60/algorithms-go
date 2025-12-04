package skiena_4

import (
	"fmt"
	"sort"
	"testing"
	"time"

	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sorting"
	"github.com/mikejlong60/algorithms-go/chapter5"
	log "github.com/sirupsen/logrus"
)

// Interesting fact about my quicksort versus golang one versus my non-parallel merge sort.  My Quicksort implementation is faster than Golang's
// as the number of duplicate elements decreases. Golang's sort gets faster the more duplication exists.
// Merge sort is always about 7x slower than either.
func TestMergeSortPerformanceVersusYourQuicksortAndGolangQuicksort(t *testing.T) {

	lt := func(l, r int) bool {
		if l < r {
			return true
		} else {
			return false
		}
	}
	eq := func(l, r int) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}
	g0 := propcheck.ChooseInt(1, 2)
	g1 := propcheck.ChooseArray(10000, 10000, g0)
	rng := propcheck.SimpleRNG{Seed: time.Now().Nanosecond()}
	prop := propcheck.ForAll(g1,
		"Test Mergesort  \n",
		func(xs []int) []int {
			return xs
		},
		func(xs []int) (bool, error) {
			var errors error

			var start = time.Now()
			l := len(xs)
			actual := chapter5.MergeSort(xs, lt)
			log.Infof("Mergesort array of:%v ints took:%v", l, time.Since(start))

			expected := make([]int, l)
			copy(expected, xs)
			start = time.Now()
			sort.Ints(expected)
			log.Infof("Golang sort array of:%v ints took:%v", l, time.Since(start))

			start = time.Now()
			gs2 := make([]int, l)
			copy(gs2, xs)
			sorting.QuickSort(gs2, lt)
			log.Infof("My quicksort array of:%v ints took:%v", l, time.Since(start))

			if !arrays.ArrayEquality(actual, expected, eq) {
				errors = fmt.Errorf("Actual:%v Expected:%v", actual, expected)
			}

			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{1, rng})
	propcheck.ExpectSuccess[[]int](t, result)

}
