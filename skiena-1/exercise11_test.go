package skiena_1

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"testing"
	"time"
)

func gcdEuclideanRecursive(a, b int) int {
	if b != 0 {
		a = gcdEuclideanRecursive(b, a%b)
	}
	return a
}

func TestEuclidGCD(t *testing.T) {
	iterativeGcdEuclidean := func(a, b int) int {
		for a != b {
			if a > b {
				a -= b
			} else {
				b -= a
			}
		}
		return a
	}

	g0 := propcheck.ChooseArray(2, 3, propcheck.ChooseInt(0, 3000))
	now := time.Now().Nanosecond()
	rng := propcheck.SimpleRNG{now}
	prop := propcheck.ForAll(g0,
		"Validate GCD evaluation  \n",
		func(xs []int) []int {
			r := gcdEuclideanRecursive(xs[0], xs[1])
			return append(xs, r)
		},
		func(xs []int) (bool, error) {

			var errors error

			expected := iterativeGcdEuclidean(xs[0], xs[1]) //aa := big.randInt(r, aSize)
			actual := xs[2]
			if actual != expected {
				errors = fmt.Errorf("Actual:%v Expected:%v", actual, expected)
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
