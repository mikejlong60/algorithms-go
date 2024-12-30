package skiena_1

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"testing"
	"time"
)

// Problem - perform integer division without using either the * or / operations.
func divide(dividend, divisor, sumSoFar, quotient int) int {
	if sumSoFar == dividend {
		return quotient
	} else if sumSoFar > dividend {
		return quotient - 1
	} else {
		sumSoFar = sumSoFar + divisor
		return divide(dividend, divisor, sumSoFar, quotient+1)
	}
}

func TestDivision(t *testing.T) {
	dividend := propcheck.ChooseInt(20000, 30000)
	divisor := propcheck.ChooseInt(1, 19999)
	g2 := propcheck.Map2(dividend, divisor, func(x, y int) []int {
		return []int{x, y}
	})
	now := time.Now().Nanosecond()
	rng := propcheck.SimpleRNG{now}
	prop := propcheck.ForAll(g2,
		"Validate integer division without * or / operators  \n",
		func(xs []int) []int {
			return xs
		},
		func(xs []int) (bool, error) {
			var errors error
			expected := xs[0] / xs[1]
			actual := divide(xs[0], xs[1], 0, 0)
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
