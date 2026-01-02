package skiena_4

import (
	"fmt"
	"testing"
	"time"

	"github.com/greymatter-io/golangz/propcheck"
)

// Given an array of consisting of ones and zeros,
// 1. Devise an algorithm that sorts the array in n-1 steps using these three comparisons: x < y, x == y, x > y
// 2. Devise another algorithm that sorts the array in 2n/3 comparisons.

func TestNMinus1Comparisons(t *testing.T) {
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	res := propcheck.ChooseArray(10, 10, propcheck.ChooseInt(0, 2))
	sortIt := func(xs []int) propcheck.Pair[[]int, int] {
		fmt.Printf("Generated array of length:%v\n", len(xs))
		totalSteps := 0
		i := 0
		maybeNextZero := 0
		for i < len(xs)-1 {
			if xs[i] < xs[i+1] {
				totalSteps++
				i++
				//maybeNextZero++
			} else if xs[i] > xs[i+1] {
				xs[i], xs[i+1] = xs[i+1], xs[i]
				//maybeNextZero = i + 1
				totalSteps++
				i++
			} else if xs[i] == xs[i+1] && xs[i] == 1 {
				//look for next 0
				for j := maybeNextZero; j < len(xs); j++ {
					totalSteps++
					if xs[j] == 0 {
						//swap them
						xs[i], xs[j] = xs[j], xs[i]
						maybeNextZero = j
						break
					}
				}
				i++
			} else {
				i++
				fmt.Println("error asdfqsgqsgasgqfg")
			}
		}
		return propcheck.Pair[[]int, int]{xs, totalSteps}
	}
	verifySort := func(p propcheck.Pair[[]int, int]) (bool, error) {
		fmt.Printf("total steps:%v\n", p.B)
		return true, nil
	}
	test := propcheck.ForAll(res, "Sort an array of ones and zeros in n-1 comparisons.", sortIt, verifySort)
	propcheck.ExpectSuccess[[]int](t, test.Run(propcheck.RunParms{10, rng}))
}
