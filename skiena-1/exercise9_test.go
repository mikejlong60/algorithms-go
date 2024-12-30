package skiena_1

import (
	"fmt"
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/propcheck"
	"testing"
	"time"
)

func TestHorner(t *testing.T) {

	hornerWithFoldLeft := func(poly []int, x int) int {
		y := func(b, a int) int {
			return a + b*x
		}
		return arrays.FoldLeft(poly, 0, y)
	}
	horner := func(poly []int, x int) int {
		n := len(poly)
		p := poly[0]
		for i := 1; i < n; i++ {
			p = p*x + poly[i]
		}
		return p
	}

	ff := func(a, b, c int) []int {
		return []int{a, b, c}
	}
	fg := func(fx []int) func(rng propcheck.SimpleRNG) ([]int, propcheck.SimpleRNG) {
		r := propcheck.ArrayOfN(fx[0], propcheck.ChooseInt(fx[1], fx[2]))
		return r
	}
	g0 := propcheck.ChooseInt(0, 5000)
	g1 := propcheck.ChooseInt(1, 500)
	g2 := propcheck.ChooseInt(699, 5000)
	g4 := propcheck.Map3(g0, g1, g2, ff)
	g5 := propcheck.FlatMap(g4, fg)
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	prop := propcheck.ForAll(g5,
		"Validate polynomial evaluation  \n",
		func(poly []int) propcheck.Pair[int, int] {
			x := poly[len(poly)-1] //x value is last element in the input array
			p1 := horner(poly[0:len(poly)-1], x)
			p2 := hornerWithFoldLeft(poly[0:len(poly)-1], x)
			return propcheck.Pair[int, int]{p1, p2}
		},
		func(p propcheck.Pair[int, int]) (bool, error) {
			var errors error
			if p.A != p.B {
				errors = fmt.Errorf("Actual:%v Expected:%v", p.A, p.B)
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
