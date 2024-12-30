package skiena_1

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"testing"
	"time"
)

var c int

func mult(y, z int) int {
	if z == 0 {
		return 0
	} else {
		a := mult(c*y, z/c) + y*(z%c)
		return a
	}
}
func TestWeirdMult(t *testing.T) {
	g0 := propcheck.ArrayOfN(3, propcheck.ChooseInt(3, 300000))
	now := time.Now().Nanosecond()
	rng := propcheck.SimpleRNG{now}
	prop := propcheck.ForAll(g0,
		"Validate weird multiplication algorithm  \n",
		func(xs []int) []int {
			c = xs[2]
			r := mult(xs[0], xs[1])
			return append(xs, r)
		},
		func(xs []int) (bool, error) {
			var errors error
			expected := xs[0] * xs[1]
			actual := xs[3]
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
