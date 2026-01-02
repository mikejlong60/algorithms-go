package skiena_4

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sorting"
)

// Given an array of A[1..n] that has numbers between 1 and n(squared), but which contains at most
// log log n different numbers, show that you can sort it in less than O(n log n).

// Generates an n array of A[1..n] that has numbers between 1 and n(squared), but which contains at most
// log log n different numbers.  The array is a size between start and stopExclusive
func ArrayOfNWithBase2LogLogNDifferentValuesOrEmptyArray(start int, stopExclusive int) func(rng propcheck.SimpleRNG) ([]int, propcheck.SimpleRNG) {
	//Generate the size of the array
	g1 := propcheck.ChooseInt(start, stopExclusive)

	//make it a float
	g2 := propcheck.Map(g1, func(x int) float64 {
		return float64(x)
	})

	// Take that number and get its log2 log2(a) and generate an array of x numbers
	// with only a different randomly generated numbers.
	g3 := func(x float64) func(propcheck.SimpleRNG) ([]int, propcheck.SimpleRNG) {
		return func(rng propcheck.SimpleRNG) ([]int, propcheck.SimpleRNG) {

			r := math.Round(math.Log2(math.Log2(x)))
			var lr = rng //The ever-changing random number generator inside the loop below.  It gets returned to ensure that the test always produces the same data for a given seed.
			xx := int(x)
			rr := int(r)
			var res = make([]int, xx)
			var a int
			var j int
			var start int
			for i := 0; i < rr; i++ {
				a, lr = propcheck.ChooseInt(start, xx/rr)(lr)
				for j = 0; j < xx/rr; j++ {
					res[start] = a
					start++
				}
			}
			//Randomize the ordering of the array
			sorting.FisherYatesShuffle(res)

			return res, lr
		}
	}
	r := propcheck.FlatMap(g2, g3)

	emptyArray := propcheck.Id[[]int]([]int{})

	//You can change the weights to produce periodic empty arrays.
	l1 := []propcheck.WeightedGen[[]int]{
		{Gen: r, Weight: 10}, {Gen: emptyArray, Weight: 1},
	}

	r2 := propcheck.Weighted(l1)
	return r2
}

func log2log2Nsort(xs []int) propcheck.Pair[[]int, []int] {
	allDistinctValues := make(map[int]int)

	for _, x := range xs {
		allDistinctValues[x] = allDistinctValues[x] + 1
	}

	keys := make([]propcheck.Pair[int, int], 0)
	for key, value := range allDistinctValues {
		keys = append(keys, propcheck.Pair[int, int]{key, value})
	}

	lt := func(l, r propcheck.Pair[int, int]) bool {
		if l.A < r.A {
			return true
		} else {
			return false
		}
	}
	sorting.QuickSort(keys, lt)

	result := make([]int, len(xs))
	var h = 0
	for _, key := range keys {
		for i := 0; i < key.B; i++ {
			result[h] = key.A
			h++
		}
	}

	return propcheck.Pair[[]int, []int]{xs, result}
}

func TestGeneratorForArrayOfNWithBase2LogLogNDifferentValues(t *testing.T) {
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	res := ArrayOfNWithBase2LogLogNDifferentValuesOrEmptyArray(10000, 100000)
	checker := func(xs []int) propcheck.Pair[[]int, []int] {
		fmt.Printf("Generated array of length:%v\n", len(xs))

		start := time.Now()
		xss := log2log2Nsort(xs)

		lt := func(l, r int) bool {
			if l < r {
				return true
			} else {
				return false
			}
		}
		//xss.B is the array sorted with my sort, xss.A is the original array.
		fmt.Printf("my sort took:%v\n", time.Since(start))
		start = time.Now()
		sorting.QuickSort(xss.A, lt)
		fmt.Printf("regular quicksort sort took:%v\n", time.Since(start))

		return xss
	}
	assertion := func(xss propcheck.Pair[[]int, []int]) (bool, error) {
		eq := func(l, r int) bool {
			if l == r {
				return true
			} else {
				return false
			}
		}

		//xss.B is the array sorted with my sort, xss.A is the original array sorted with quicksort.
		equal := arrays.ArrayEquality(xss.A, xss.B, eq)
		if !equal {
			return false, fmt.Errorf("arrays not equal")
		}
		return true, nil
	}
	test := propcheck.ForAll(res, "Weighted should have produced A number between 1000 and 5000 exclusive or between 100000 and 200000 exclusive.", checker, assertion)
	propcheck.ExpectSuccess[[]int](t, test.Run(propcheck.RunParms{10, rng}))
}
