package skiena_2

import (
	"fmt"
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
	"testing"
	"time"
)

func TestEqualProbabilitySubset(t *testing.T) {
	eq := func(l, r int) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}
	lt := func(l, r int) bool {
		if l < r {
			return true
		} else {
			return false
		}
	}
	//Make an array of k numbers
	f := propcheck.ChooseArray(0, 20, propcheck.ChooseInt(1, 1000000))

	//Randomly choose the sub-array of numbers xs from set k with a random size and equal probability of each member of xs.
	//Note this is a pure function
	g := func(set []int) func(r propcheck.SimpleRNG) (propcheck.Pair[[]int, []int], propcheck.SimpleRNG) {
		a := propcheck.ChooseArray(0, len(set), propcheck.ChooseInt(0, len(set)))

		i := func(subsetIdx []int) propcheck.Pair[[]int, []int] {
			r := []int{}
			for i := 0; i < len(subsetIdx); i++ {
				r = append(r, set[subsetIdx[i]])
			}

			r2 := sets.ToSet(r, lt, eq)
			set2 := sets.ToSet(set, lt, eq)
			return propcheck.Pair[[]int, []int]{r2, set2}
		}

		return propcheck.Map(a, i)
	}

	h := propcheck.FlatMap(f, g)
	now := time.Now().Nanosecond()
	rng := propcheck.SimpleRNG{now}

	passThru := func(a propcheck.Pair[[]int, []int]) propcheck.Pair[[]int, []int] {
		return a
	}

	subSetHasOnlyElementsOfSet := propcheck.ForAll(h, fmt.Sprintf("Subset must only have elements of set \n"),
		passThru,
		func(a propcheck.Pair[[]int, []int]) (bool, error) {
			var errors error
			subset := a.A
			set := a.B
			if len(set) > 0 && !arrays.ContainsAllOf(set, subset, eq) {
				errors = fmt.Errorf("Set:%v did not contain all of subset:%v", set, subset)
			}
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)

	subSetIsNoLargerThanSet := propcheck.ForAll(h, fmt.Sprintf("Subset cannot be larger than set \n"),
		passThru,
		func(a propcheck.Pair[[]int, []int]) (bool, error) {
			var errors error
			subset := a.A
			set := a.B
			if len(subset) > len(set) {
				errors = fmt.Errorf("Subset:%v was larger than set:%v", subset, set)
			}
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)

	test := propcheck.And[propcheck.Pair[[]int, []int]](subSetHasOnlyElementsOfSet, subSetIsNoLargerThanSet)
	result := test.Run(propcheck.RunParms{400, rng})
	propcheck.ExpectSuccess[propcheck.Pair[[]int, []int]](t, result)
}
