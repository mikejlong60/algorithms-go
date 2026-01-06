package skiena_4

import (
	"fmt"
	"testing"
	"time"

	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sorting"
)

// Given an array of consisting of ones and zeros,
// 1. Devise an algorithm that sorts the array in n-1 steps using these three comparisons: x < y, x == y, x > y
// 2. Devise another algorithm that sorts the array in 2n/3 comparisons.

func TestNMinus1Comparisons(t *testing.T) {
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	res := propcheck.ChooseArray(100000, 1000000, propcheck.ChooseInt(0, 2))
	sortIt := func(xs []int) []int {
		fmt.Printf("Generated array of length:%v\n", len(xs))
		i := 0
		maybeNextZero := 0
		for i < len(xs)-1 {
			if xs[i] == 1 { //see if there are any zeros past here
				for j := maybeNextZero; j < len(xs); j++ {
					if xs[j] == 0 {
						xs[i], xs[j] = xs[j], xs[i]
						maybeNextZero = j + 1
						break
					}
					if j == len(xs)-1 {
						return xs
					}
				}
			} else {
				maybeNextZero = i + 1
			}
			i++
		}
		return xs
	}
	verifySort := func(actual []int) (bool, error) {
		//fmt.Printf("xs:%v\n", actual)
		lt := func(x, y int) bool {
			return x < y
		}
		expected := make([]int, len(actual))
		copy(expected, actual)
		sorting.QuickSort(expected, lt)
		r := arrays.ArrayEquality(actual, expected, func(l, r int) bool { return l == r })
		if !r {
			return false, fmt.Errorf("expected %v, got %v", expected, actual)
		}
		return true, nil
	}
	test := propcheck.ForAll(res, "Sort an array of ones and zeros in n-1 comparisons.", sortIt, verifySort)
	propcheck.ExpectSuccess[[]int](t, test.Run(propcheck.RunParms{10, rng}))
}
