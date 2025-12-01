package skiena_4

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
)

//Implement wiggle sort.  The definition of wiggle sort is make an unsorted array of integers
//ordered such that a[0] < a[1] and a[2] < a[1] and a[2] < a[3]

// Do it in O(n) in space O(1)
//Possible answers for a == [3,1,4,2,6,5] are [1,3,2,5,4,6]

//Possible answers for a == [3,1,4,2,6,5, 7, 8] are [1,3,2,5,4,7,6,8]

//1,3,2,4,6,5

func wiggleSort(xs []int) []int {
	for i := 0; i < len(xs); {
		if i+2 < len(xs) && xs[i] > xs[i+1] {
			xs[i], xs[i+1] = xs[i+1], xs[i]
		}
		i = i + 2
	}

	for i := 1; i < len(xs); {
		if i+2 < len(xs) && xs[i] < xs[i+1] {
			xs[i], xs[i+1] = xs[i+1], xs[i]
		}
		i = i + 2
	}
	return xs
}

func TestWiggleSort(t *testing.T) {
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

	maxListSize := 3000
	minListSize := 10
	ge := propcheck.ChooseInt(0, 20000)
	ge2 := sets.ChooseSet(minListSize, maxListSize, ge, lt, eq)
	rng := propcheck.SimpleRNG{Seed: time.Now().Nanosecond()}
	runWiggleSort := func(xs []int) []int {
		var r []int //DEMONSTRATE THAT THIS IS A PURE FUNCTION, INVOKED FOR EACH ANDED RESULT WITH EACH SET TEST SET.
		if (len(xs) % 2) != 0 {
			r = xs[1:]
		} else {
			r = xs
		}
		rand.Shuffle(len(r), func(i, j int) {
			r[i], r[j] = r[j], r[i]
		})
		wiggleSort(r)
		return r
	}

	evenElementsAscend := propcheck.ForAll(ge2, "even element [i] < [i+1]",
		runWiggleSort,
		func(xs []int) (bool, error) {
			for i := 0; i < len(xs); {
				if i+2 < len(xs) && xs[i] > xs[i+1] {
					return false, fmt.Errorf("even element [i] > [i+1]")
				}
				i = i + 2
			}
			return true, nil
		},
	)
	oddElementsAscend := propcheck.ForAll(ge2, "odd element [i] > [i+1]",
		runWiggleSort,
		func(xs []int) (bool, error) {
			for i := 1; i < len(xs); {
				if i+2 < len(xs) && xs[i] < xs[i+1] {
					return false, fmt.Errorf("odd element [i] > [i+1]")
				}
				i = i + 2
			}
			return true, nil
		},
	)
	bigProp := propcheck.And[[]int](evenElementsAscend, oddElementsAscend)
	result := bigProp.Run(propcheck.RunParms{1, rng})
	propcheck.ExpectSuccess[[]int](t, result)
}
