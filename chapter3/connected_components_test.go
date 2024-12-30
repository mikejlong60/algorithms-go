package chapter3

import (
	"fmt"
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
	"github.com/hashicorp/go-multierror"
	"testing"
	"time"
)

func TestAllConnectedComponents(t *testing.T) {
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
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

	h := func(stuffToAssert propcheck.Pair[[][]Edge, []int]) (bool, error) {
		g := func(connectedComponent []Edge) []int {
			var r = []int{}
			for _, i := range connectedComponent {
				r = append(r, i.U)
				r = append(r, i.V)
			}
			return sets.ToSet(r, lt, eq)
		}
		allConnectedNodes := arrays.FlatMap(stuffToAssert.A, g)
		var errors error
		if len(allConnectedNodes) != len(stuffToAssert.B) {
			errors = multierror.Append(errors, fmt.Errorf("Number of Nodes:%v in set of connected components should have equaled same number of Nodes:%v", len(allConnectedNodes), len(stuffToAssert.B)))
		}
		if !sets.SetEquality(allConnectedNodes, stuffToAssert.B, eq) {
			errors = multierror.Append(errors, fmt.Errorf("Not every node was connected"))
		}
		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}

	prop := propcheck.ForAll(UndirectedGraphGen(1, 100),
		"Generate the set of connected components for a given graph.",
		GenerateConnectedComponents,
		h,
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[propcheck.Pair[map[int]*Node, int]](t, result)
}
