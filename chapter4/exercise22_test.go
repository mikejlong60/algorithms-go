package chapter4

import (
	"github.com/greymatter-io/golangz/sets"
	"testing"
)

// This test informally proves that the lowest cost edge, cb, is in every MST.
// And I can make new same-weight MST by altering the order in the input array of edges.
// And every MST in the example here has the same cost.
func TestLowestCostEdgeInEveryMSTWithNonDistinctValues(t *testing.T) {
	ab := &PrimsEdge{
		u:      "a",
		v:      "b",
		weight: 4,
	}

	ad := &PrimsEdge{
		u:      "a",
		v:      "d",
		weight: 4,
	}

	ac := &PrimsEdge{
		u:      "a",
		v:      "c",
		weight: 4,
	}

	bd := &PrimsEdge{
		u:      "b",
		v:      "d",
		weight: 4,
	}

	cd := &PrimsEdge{
		u:      "c",
		v:      "d",
		weight: 4,
	}
	lowestCostEdge := &PrimsEdge{
		u:      "c",
		v:      "b",
		weight: 3,
	}

	eq := func(l, r *PrimsEdge) bool {
		if l.u == r.u && l.v == r.v && l.weight == r.weight {
			return true
		} else {
			return false
		}
	}
	var actual = Kruskals([]*PrimsEdge{ad, ab, ac, bd, cd, lowestCostEdge})
	var expected = []*PrimsEdge{ac, lowestCostEdge, cd}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}

	//Alter the order of the input which results in a different but same cost MST that includes the lowest cost edge
	actual = Kruskals([]*PrimsEdge{ad, bd, cd, lowestCostEdge, ab, ac})
	expected = []*PrimsEdge{ac, lowestCostEdge, bd}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}

	//Alter the order of the input which results in a different but same cost MST that includes the lowest cost edge
	actual = Kruskals([]*PrimsEdge{ad, lowestCostEdge, ab, bd, cd, ac})
	expected = []*PrimsEdge{ac, lowestCostEdge, cd}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}

	//Alter the order of the input which results in a different but same cost MST that includes the lowest cost edge
	actual = Kruskals([]*PrimsEdge{ad, ab, bd, cd, ac, lowestCostEdge})
	expected = []*PrimsEdge{ac, lowestCostEdge, bd}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}

}
