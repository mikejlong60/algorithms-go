package chapter4

import (
	"github.com/greymatter-io/golangz/sets"
	"testing"
)

// The THe following two tests prove that a graph with duplicate edges has more than 1 MST and the Kruskals algorithm still
// produces a MST in such a case depending on the ordering of the graph.
func TestDupEdgeWeightProducesValidTreeKruskals(t *testing.T) {
	ab := &PrimsEdge{
		u:      "a",
		v:      "b",
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
	cb := &PrimsEdge{
		u:      "c",
		v:      "b",
		weight: 4,
	}

	eq := func(l, r *PrimsEdge) bool {
		if l.u == r.u && l.v == r.v && l.weight == r.weight {
			return true
		} else {
			return false
		}
	}
	//Try several orderings to make sure you get different MST with same total weight
	var actual = Kruskals([]*PrimsEdge{ab, ac, bd, cd, cb})
	var expected = []*PrimsEdge{ac, cd, ab}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}

	actual = Kruskals([]*PrimsEdge{bd, cd, cb, ab, ac})
	expected = []*PrimsEdge{ac, bd, ab}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}

}

func TestDupEdgeWeightProducesValidTreeKruskalsUsingUnionFind(t *testing.T) {
	ab := &PrimsEdge{
		u:      "a",
		v:      "b",
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
	cb := &PrimsEdge{
		u:      "c",
		v:      "b",
		weight: 4,
	}

	eq := func(l, r *PrimsEdge) bool {
		if l.u == r.u && l.v == r.v && l.weight == r.weight {
			return true
		} else {
			return false
		}
	}
	//Try several orderings to make sure you get different MST with same total weight
	var actual = KruskalsUsingUnionFind([]*PrimsEdge{ab, ac, bd, cd, cb})
	var expected = []*PrimsEdge{ac, cb, bd}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}

	actual = KruskalsUsingUnionFind([]*PrimsEdge{bd, cd, cb, ab, ac})
	expected = []*PrimsEdge{ac, cb, cd}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}

}
