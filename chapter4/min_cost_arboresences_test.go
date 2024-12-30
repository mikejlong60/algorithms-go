package chapter4

import (
	"github.com/greymatter-io/golangz/sets"
	"testing"
)

func TestExampleFromVideoWithNoCycles(t *testing.T) {

	eq := func(l, r Edge) bool {
		if l.u == r.u && l.v == r.v {
			return true
		} else {
			return false
		}
	}

	a := Node{id: "a"}
	b := Node{id: "b"}
	c := Node{id: "c"}
	d := Node{id: "d"}
	e := Node{id: "e"}
	f := Node{id: "f"}

	ab := Edge{&a, &b, 2}
	ac := Edge{&a, &c, 10}
	ad := Edge{&a, &d, 2}
	be := Edge{&b, &e, 2}
	bf := Edge{&b, &f, 8}
	cf := Edge{&c, &f, 8}
	cd := Edge{&c, &d, 2}
	db := Edge{&d, &b, 4}
	ec := Edge{&e, &c, 2}

	b.nodesEntering = []Edge{db, ab}
	c.nodesEntering = []Edge{ac, ec}
	d.nodesEntering = []Edge{cd, ad}
	e.nodesEntering = []Edge{be}
	f.nodesEntering = []Edge{cf, bf}

	s := []*Node{&a, &b, &c, &d, &e, &f}
	actual := MinCost(s, &a)
	expected := []Edge{ab, be, cd, ec, cf}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}

func TestExampleFromVideoWithOneCycle(t *testing.T) {

	eq := func(l, r Edge) bool {
		if l.u == r.u && l.v == r.v {
			return true
		} else {
			return false
		}
	}

	a := Node{id: "a"}
	b := Node{id: "b"}
	c := Node{id: "c"}
	d := Node{id: "d"}
	e := Node{id: "e"}
	f := Node{id: "f"}

	ab := Edge{&a, &b, 10}
	ac := Edge{&a, &c, 10}
	ad := Edge{&a, &d, 2}
	be := Edge{&b, &e, 2}
	bf := Edge{&b, &f, 4}
	cf := Edge{&c, &f, 8}
	cd := Edge{&c, &d, 1}
	db := Edge{&d, &b, 4}
	ec := Edge{&e, &c, 2}

	b.nodesEntering = []Edge{db, ab}
	c.nodesEntering = []Edge{ac, ec}
	d.nodesEntering = []Edge{cd, ad}
	e.nodesEntering = []Edge{be}
	f.nodesEntering = []Edge{cf, bf}

	s := []*Node{&a, &b, &c, &d, &e, &f}
	actual := MinCost(s, &a)
	expected := []Edge{ad, db, bf, be, ec}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", actual, expected)
	}
}
