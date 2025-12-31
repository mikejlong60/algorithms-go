package skiena_4

import (
	"fmt"
	"math"
	"testing"

	"github.com/greymatter-io/golangz/propcheck"
)

// Given an array of A[1..n] that has numbers between 1 and n(squared), but which contains at most
// log log n different numbers, show that you can sort it in less than O(n log n).

// Generates an n array of A[1..n] that has numbers between 1 and n(squared), but which contains at most
// log log n different numbers.  The array is a size between start and stopExclusive
func ArrayOfNWithBase2LogLogNDifferentValuesOrEmptyArray(start int, stopExclusive int) func(rng propcheck.SimpleRNG) ([]int, propcheck.SimpleRNG) {
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
			return res, lr
		}
	}
	r := propcheck.FlatMap(g2, g3)

	emptyArray := propcheck.Id[[]int]([]int{})

	//You can change the weights to produce periodic empty arrays.
	l1 := []propcheck.WeightedGen[[]int]{
		{Gen: r, Weight: 10}, {Gen: emptyArray, Weight: 10},
	}

	r2 := propcheck.Weighted(l1)
	return r2
}

func TestGeneratorForArrayOfNWithBase2LogLogNDifferentValues(t *testing.T) {
	rng := propcheck.SimpleRNG{86322638} //time.Now().Nanosecond()}
	res := ArrayOfNWithBase2LogLogNDifferentValuesOrEmptyArray(100, 1000)
	checker := func(xs []int) []int {
		fmt.Printf("Generated array of length:%v\n", len(xs))

		return xs
	}
	assertion := func(xs []int) (bool, error) {
		if len(xs) < 10 {
			return false, fmt.Errorf("not enough elements:%v", len(xs))
		}
		return true, nil
	}
	test := propcheck.ForAll(res, "Weighted should have produced A number between 1000 and 5000 exclusive or between 100000 and 200000 exclusive.", checker, assertion)
	propcheck.ExpectSuccess[[]int](t, test.Run(propcheck.RunParms{10, rng}))
}
