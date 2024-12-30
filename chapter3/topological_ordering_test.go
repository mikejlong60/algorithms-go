package chapter3

import (
	"fmt"
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/hashicorp/go-multierror"
	"testing"
	"time"
)

func TestTopologicalOrderingFromBookExercise37(t *testing.T) {
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}

	makeExcercise37Graph := func(a propcheck.Pair[map[int]*Node, int]) map[int]*NodeForTopoOrdering {

		v1 := &NodeForTopoOrdering{
			Id:                  1,
			OutgoingConnections: make(map[int]*NodeForTopoOrdering),
			IncomingConnections: make(map[int]*NodeForTopoOrdering),
		}
		v2 := &NodeForTopoOrdering{
			Id:                  2,
			OutgoingConnections: make(map[int]*NodeForTopoOrdering),
			IncomingConnections: make(map[int]*NodeForTopoOrdering),
		}
		v3 := &NodeForTopoOrdering{
			Id:                  3,
			OutgoingConnections: make(map[int]*NodeForTopoOrdering),
			IncomingConnections: make(map[int]*NodeForTopoOrdering),
		}
		v4 := &NodeForTopoOrdering{
			Id:                  4,
			OutgoingConnections: make(map[int]*NodeForTopoOrdering),
			IncomingConnections: make(map[int]*NodeForTopoOrdering),
		}
		v5 := &NodeForTopoOrdering{
			Id:                  5,
			OutgoingConnections: make(map[int]*NodeForTopoOrdering),
			IncomingConnections: make(map[int]*NodeForTopoOrdering),
		}
		v6 := &NodeForTopoOrdering{
			Id:                  6,
			OutgoingConnections: make(map[int]*NodeForTopoOrdering),
			IncomingConnections: make(map[int]*NodeForTopoOrdering),
		}
		v7 := &NodeForTopoOrdering{
			Id:                  7,
			OutgoingConnections: make(map[int]*NodeForTopoOrdering),
			IncomingConnections: make(map[int]*NodeForTopoOrdering),
		}

		v1.OutgoingConnections[5] = v5
		v1.OutgoingConnections[7] = v7
		v1.OutgoingConnections[4] = v4

		v2.OutgoingConnections[6] = v6
		v2.OutgoingConnections[5] = v5
		v2.OutgoingConnections[3] = v3

		v3.OutgoingConnections[4] = v4
		v3.OutgoingConnections[5] = v5
		v3.IncomingConnections[2] = v2

		v4.OutgoingConnections[5] = v5
		v4.IncomingConnections[1] = v1
		v4.IncomingConnections[3] = v3

		v5.OutgoingConnections[6] = v6
		v5.OutgoingConnections[7] = v7
		v5.IncomingConnections[1] = v1
		v5.IncomingConnections[2] = v2
		v5.IncomingConnections[3] = v3
		v5.IncomingConnections[4] = v4

		v6.OutgoingConnections[7] = v7
		v6.IncomingConnections[5] = v5
		v6.IncomingConnections[2] = v2

		v7.IncomingConnections[6] = v6
		v7.IncomingConnections[5] = v5
		v7.IncomingConnections[1] = v1

		connectedComponents := make(map[int]*NodeForTopoOrdering)
		connectedComponents[1] = v1
		connectedComponents[2] = v2
		connectedComponents[3] = v3
		connectedComponents[4] = v4
		connectedComponents[5] = v5
		connectedComponents[6] = v6
		connectedComponents[7] = v7

		return connectedComponents
	}

	makeExpectedTopo := func() []*NodeForTopoOrdering {
		v1 := &NodeForTopoOrdering{
			Id:                  1,
			OutgoingConnections: make(map[int]*NodeForTopoOrdering),
			IncomingConnections: make(map[int]*NodeForTopoOrdering),
		}
		v2 := &NodeForTopoOrdering{
			Id:                  2,
			OutgoingConnections: make(map[int]*NodeForTopoOrdering),
			IncomingConnections: make(map[int]*NodeForTopoOrdering),
		}
		v3 := &NodeForTopoOrdering{
			Id:                  3,
			OutgoingConnections: make(map[int]*NodeForTopoOrdering),
			IncomingConnections: make(map[int]*NodeForTopoOrdering),
		}
		v4 := &NodeForTopoOrdering{
			Id:                  4,
			OutgoingConnections: make(map[int]*NodeForTopoOrdering),
			IncomingConnections: make(map[int]*NodeForTopoOrdering),
		}
		v5 := &NodeForTopoOrdering{
			Id:                  5,
			OutgoingConnections: make(map[int]*NodeForTopoOrdering),
			IncomingConnections: make(map[int]*NodeForTopoOrdering),
		}
		v6 := &NodeForTopoOrdering{
			Id:                  6,
			OutgoingConnections: make(map[int]*NodeForTopoOrdering),
			IncomingConnections: make(map[int]*NodeForTopoOrdering),
		}
		v7 := &NodeForTopoOrdering{
			Id:                  7,
			OutgoingConnections: make(map[int]*NodeForTopoOrdering),
			IncomingConnections: make(map[int]*NodeForTopoOrdering),
		}

		v1.OutgoingConnections[5] = v5
		v1.OutgoingConnections[7] = v7
		v1.OutgoingConnections[4] = v4

		v2.OutgoingConnections[6] = v6
		v2.OutgoingConnections[5] = v5
		v2.OutgoingConnections[3] = v3

		v3.OutgoingConnections[4] = v4
		v3.OutgoingConnections[5] = v5

		v4.OutgoingConnections[5] = v5

		v5.OutgoingConnections[6] = v6
		v5.OutgoingConnections[7] = v7

		v6.OutgoingConnections[7] = v7

		return []*NodeForTopoOrdering{v1, v2, v3, v4, v5, v6, v7}
	}

	eq := func(a, b *NodeForTopoOrdering) bool {
		nodeEq := func(aa, bb int) bool {
			if aa == bb {
				return true
			} else {
				return false
			}
		}

		getKeys := func(a map[int]*NodeForTopoOrdering) []int {
			keys := make([]int, len(a))
			i := 0
			for k := range a {
				keys[i] = k
				i++
			}
			return keys
		}

		x := arrays.ArrayEquality(getKeys(a.OutgoingConnections), getKeys(b.OutgoingConnections), nodeEq)
		fmt.Println(x)
		if a.Id == b.Id &&
			arrays.ArrayEquality(getKeys(a.IncomingConnections), getKeys(b.IncomingConnections), nodeEq) { //&&
			//arrays.ArrayEquality(getKeys(a.OutgoingConnections), getKeys(b.OutgoingConnections), nodeEq) {
			return true
		} else {
			return false
		}
	}

	prop := propcheck.ForAll(UndirectedGraphGen(1, 100),
		"Verify the topological ordering from figure 3.7 in Algorithms Book",
		makeExcercise37Graph,
		func(xs map[int]*NodeForTopoOrdering) (bool, error) {
			_, actual := Topo(xs, []*NodeForTopoOrdering{})
			expected := makeExpectedTopo()
			var errors error
			if !arrays.ArrayEquality(actual, expected, eq) {
				errors = multierror.Append(errors, fmt.Errorf("expected topo:%v, actual topo:%v", expected, actual))
			}
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{1, rng})
	propcheck.ExpectSuccess[propcheck.Pair[map[int]*Node, int]](t, result)
}
