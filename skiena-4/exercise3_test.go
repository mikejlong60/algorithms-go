package skiena_4

import (
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/mikejlong60/algorithms/chapter5"
	"math"
	"testing"
)

func TestMinimumMaximumSum(t *testing.T) {
	//Big O(n log n) because of sorting.
	lowestMaximumPairSum := func(xss []int) []propcheck.Pair[int, int] {
		lt := func(l, r int) bool {
			if l < r {
				return true
			} else {
				return false
			}
		}
		rr := chapter5.MergeSort(xss, lt)
		actual := make([]propcheck.Pair[int, int], 0)
		for i := 0; i < len(rr)/2; i++ {
			pair := propcheck.Pair[int, int]{rr[i], rr[(len(rr)-1)-i]}
			actual = append(actual, pair)
		}
		return actual
	}

	g := func(min int, p propcheck.Pair[int, int]) int {
		if p.A+p.B > min {
			return p.A + p.B
		} else {
			return min
		}
	}
	actual := lowestMaximumPairSum([]int{1, 3, 5, 9})

	largestPairSum := arrays.FoldLeft(actual, math.MinInt, g)
	if largestPairSum != 10 {
		t.Errorf("\nActual largestPairSum:  %v\nExpected largestPairSum:%v", largestPairSum, 10)
	}
	actual = lowestMaximumPairSum([]int{1, 1, 13, 14, 233, 232, 234, 43, 13, 45, 15, 26, 7, 18, 90, 1, 12, 23})
	largestPairSum = arrays.FoldLeft(actual, math.MinInt, g)
	if largestPairSum != 235 {
		t.Errorf("\nActual largestPairSum:  %v\nExpected largestPairSum:%v", largestPairSum, 235)
	}
}
