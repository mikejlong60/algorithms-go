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

func TestEqualityOfNodesInDfAndBfSearch(t *testing.T) {
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}

	prop := propcheck.ForAll(UndirectedGraphGen(1, 100),
		"Generate a random graph and do a Tree search starting from some root.",
		func(graph propcheck.Pair[map[int]*Node, int]) []int {
			var tree []Edge
			dfStart := time.Now()
			_, _, dfTree := DFSearch(graph.A[graph.B], make(map[int]*Node), tree)
			fmt.Printf("DFS on a graph of size:%v took %v\n", len(graph.A), time.Since(dfStart))
			bfStart := time.Now()
			bfTree, _, _ := BFSearch(graph.A, graph.B)
			fmt.Printf("BFS on a graph of size:%v took %v\n", len(graph.A), time.Since(bfStart))
			bf := func(e []Edge) []int {
				var r []int
				for _, b := range e {
					r = append(r, b.U)
					r = append(r, b.V)
				}
				return r
			}
			df := func(e Edge) []int {
				var r []int
				r = append(r, e.U)
				r = append(r, e.V)
				return r
			}
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
			dfConnectedComponent := sets.ToSet(arrays.FlatMap(dfTree, df), lt, eq)
			bfConnectedComponent := sets.SetMinus(sets.ToSet(arrays.FlatMap(bfTree, bf), lt, eq), []int{-1}, eq) //Remove the default -1 first node
			if len(bfConnectedComponent) == 1 && len(dfConnectedComponent) == 0 {                                //If root node has no connections then df has no tree and this is fine.
				return []int{}
			} else {
				return sets.SetMinus(bfConnectedComponent, dfConnectedComponent, eq)
			}
		},
		func(xs []int) (bool, error) {
			var errors error
			if len(xs) > 0 {
				errors = multierror.Append(errors, fmt.Errorf("Depth-first and breadth-first search should have produced the same set of components but differed(bf - df=%v\n", xs))
			}
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[propcheck.Pair[map[int]*Node, int]](t, result)
}
