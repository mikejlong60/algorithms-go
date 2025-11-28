package skiena_4

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/greymatter-io/golangz/arrays"
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
	//if !(xs[0] < xs[1] && xs[2] < xs[1] && xs[2] < xs[3]) {
	//	xs[i], xs[3] = xs[3], xs[i]
	//}
	//if i < len(xs)-2 && xs[i] > xs[i+1] && xs[i] < xs[i+2] { //1,3,2,4 case
	//	xs[i], xs[i+1] = xs[i+1], xs[i]
	//	continue
	//}
	//if i < len(xs)-3 && xs[i] > xs[i+1] && xs[i] > xs[i+2] { //4,2,6,5 case
	//	//swap 2[i+1],4[i] then you have 2,4,6,5, then swap 4[i+1] and 5[i+3]
	//	xs[i+1], xs[i] = xs[i], xs[i+1]
	//	xs[i+1], xs[i+3] = xs[i+3], xs[i+1]
	//	continue
	//}
	//if i < len(xs)-1 && xs[i] > xs[i+1] { // now you can fall through
	//	xs[i], xs[i+1] = xs[i+1], xs[i]
	//	continue
	//}
	//}

	return xs
}

func TestWiggleSort2(t *testing.T) {
	xs := []int{3, 1, 4, 2, 6, 5, 7, 8}

	actual := wiggleSort(xs)
	expected := []int{1, 3, 2, 5, 4, 7, 6, 8}
	eq := func(l, r int) bool {
		if l == r {
			return true
		}
		return false
	}
	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
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

	maxListSize := 100 //Am banking on the odds being near zero that I get the same bool over 3000 rolls of the die.
	minListSize := 10
	ge := propcheck.ChooseInt(0, 20000)
	ge2 := sets.ChooseSet(minListSize, maxListSize, ge, lt, eq)
	rng := propcheck.SimpleRNG{Seed: time.Now().Nanosecond()}
	runWiggleSort := func(xs []int) []int {
		var r []int
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

	evenElementsAscend := propcheck.ForAll(ge2, "even element pairs ascend",
		runWiggleSort,
		func(xs []int) (bool, error) {
			for i := 0; i < len(xs); {
				if i+2 < len(xs) && xs[i] > xs[i+1] {
					return false, fmt.Errorf("even integers not in ascending sequence")
				}
				i = i + 2
			}
			return true, nil
		},
	)
	oddElementsAscend := propcheck.ForAll(ge2, "uneven element pairs descend",
		runWiggleSort,
		func(xs []int) (bool, error) {
			for i := 1; i < len(xs); {
				if i+2 < len(xs) && xs[i] < xs[i+1] {
					return false, fmt.Errorf("uneven integers not in descending sequence")
				}
				i = i + 2
			}
			return true, nil
		},
	)
	bigProp := propcheck.And[[]int](evenElementsAscend, oddElementsAscend)
	result := bigProp.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]int](t, result)
}
