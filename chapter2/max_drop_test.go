package chapter2

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
	"github.com/hashicorp/go-multierror"
	"testing"
	"time"
)

func TestMaxDropGivenMaxBrokenJars(t *testing.T) { //This reverses the order because foldRight does that. There is a Reverse function in Golangz if that matters to you.
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

	g0 := propcheck.ChooseInt(0, 10000000)
	g1 := sets.ChooseSet(0, 500000, g0, lt, eq) //This array comes back sorted

	f := func(xss []int) func(propcheck.SimpleRNG) (propcheck.Pair[int, propcheck.Pair[[]int, int]], propcheck.SimpleRNG) {
		g := propcheck.ChooseInt(0, len(xss)-1)  //breakingPoint index
		gg := propcheck.ChooseInt(2, len(xss)/2) //Maximum Number of broken jars
		i := func(x, maxBrokenJars int) propcheck.Pair[int, propcheck.Pair[[]int, int]] {
			ladderAndBreakingPoint := propcheck.Pair[[]int, int]{
				A: xss,
				B: xss[x],
			}
			r := propcheck.Pair[int, propcheck.Pair[[]int, int]]{
				A: maxBrokenJars,
				B: ladderAndBreakingPoint,
			}
			return r
		}
		h := propcheck.Map2(g, gg, i)
		return h
	}

	g2 := propcheck.FlatMap(g1, f)
	now := time.Now().Nanosecond()
	rng := propcheck.SimpleRNG{now}
	prop := propcheck.ForAll(g2,
		"Exercise 2.8b, the max jar drop given a budget of not-to-exceed broken jars.",
		func(xs propcheck.Pair[int, propcheck.Pair[[]int, int]]) propcheck.Pair[int, propcheck.Pair[[]int, int]] {
			//A is the ladder, B is the breaking point on the ladder(the actual value in the array, not its index).
			numberOfSteps = 0
			numberOfSingleSteps = 0
			r := HighestBreakingPoint(xs.B.A, xs.B.A, xs.B.B, xs.A, 0)
			fmt.Printf("number of steps == numberOfSingleSteps:%v + numberOfBinarySearches:%v for array of size:%v total steps:%v\n", numberOfSingleSteps, numberOfSteps, len(xs.B.A), numberOfSingleSteps+numberOfSteps)
			return propcheck.Pair[int, propcheck.Pair[[]int, int]]{r, xs.B}
		},
		func(highestWrungAEtAll propcheck.Pair[int, propcheck.Pair[[]int, int]]) (bool, error) {
			var errors error
			breakingPoint := highestWrungAEtAll.B.B
			highestWrung := highestWrungAEtAll.A
			ladder := highestWrungAEtAll.B.A
			var highestWrungIdx = -1
			for i := 0; i < len(ladder); i++ {
				if ladder[i] == highestWrung {
					if ladder[i] == highestWrung {
						highestWrungIdx = i
						break
					}
				}
			}
			if ladder[highestWrungIdx] != highestWrung {
				errors = multierror.Append(errors, fmt.Errorf("Expected highest non-breaking wrung withing budget to be:%v but was:%v", highestWrung, breakingPoint))
			}
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[propcheck.Pair[int, propcheck.Pair[[]int, int]]](t, result)
}
