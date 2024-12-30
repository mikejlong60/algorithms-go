package chapter3

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/hashicorp/go-multierror"
	"testing"
	"time"
)

func TestHalfConnectedNodesIsConnectedGraph(t *testing.T) {
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}

	h := func(graph propcheck.Pair[[][]Edge, []int]) (bool, error) {
		var errors error
		//There should only be one connected component which proves
		//by induction that the entire graph is connected when every node
		//is connected to at least half of the other nodes.
		if len(graph.A) != 1 {
			errors = multierror.Append(errors, fmt.Errorf("should have been exactly one connected component but there were:%v", len(graph.A)))
		}

		if errors != nil {
			return false, errors
		} else {
			return true, nil
		}
	}

	prop := propcheck.ForAll(EvenNumberOfNodesGen(1, 1000),
		"Prove that a graph of n nodes where each is connected to at least half of the other nodes is connected",
		MakeConnectionComponents,
		h,
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[map[int]*Node](t, result)
}
