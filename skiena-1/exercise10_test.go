package skiena_1

import (
	"fmt"
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/propcheck"
	"sort"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	//mutates input array
	//Big O n squared -- not very efficient
	bubbleSort := func(xs []int) []int {
		for i := len(xs); i > 0; i-- {
			for j := 0; j < len(xs)-1; j++ {
				if xs[j] > xs[j+1] {
					newJ := xs[j]
					newJPlusOne := xs[j+1]
					xs[j] = newJPlusOne
					xs[j+1] = newJ
				}
			}
		}
		return xs
	}

	g0 := propcheck.ChooseArray(0, 100, propcheck.ChooseInt(0, 3000))
	now := time.Now().Nanosecond()
	rng := propcheck.SimpleRNG{now}
	prop := propcheck.ForAll(g0,
		"Validate bubble sort evaluation  \n",
		func(xs []int) propcheck.Pair[[]int, []int] {

			xsCp := make([]int, len(xs))

			copy(xsCp, xs)
			bubbleSort(xsCp)
			return propcheck.Pair[[]int, []int]{xsCp, xs}
		},
		func(xss propcheck.Pair[[]int, []int]) (bool, error) {
			eq := func(x, y int) bool {
				if x == y {
					return true
				} else {
					return false
				}
			}
			var errors error
			actual := xss.A
			sort.Ints(xss.B)
			if !arrays.ArrayEquality(xss.A, xss.B, eq) {
				errors = fmt.Errorf("Actual:%v Expected:%v", actual, xss.B)
			}
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]int](t, result)
}
