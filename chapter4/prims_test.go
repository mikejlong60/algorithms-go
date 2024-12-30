package chapter4

import (
	"fmt"
	"github.com/greymatter-io/golangz/heap"
	"github.com/greymatter-io/golangz/sets"
	"testing"
)

var edgeEq = func(l, r *PrimsEdge) bool {
	if l.weight == r.weight && l.u == r.u && l.v == r.v {
		return true
	} else {
		return false
	}
}

func TestMinSpanningTree(t *testing.T) {
	a := &PrimsNode{id: "a", connections: heap.New[PrimsEdge, string](extractor)}
	b := &PrimsNode{id: "b", connections: heap.New[PrimsEdge, string](extractor)}
	c := &PrimsNode{id: "c", connections: heap.New[PrimsEdge, string](extractor)}
	d := &PrimsNode{id: "d", connections: heap.New[PrimsEdge, string](extractor)}
	e := &PrimsNode{id: "e", connections: heap.New[PrimsEdge, string](extractor)}
	f := &PrimsNode{id: "f", connections: heap.New[PrimsEdge, string](extractor)}
	g := &PrimsNode{id: "g", connections: heap.New[PrimsEdge, string](extractor)}
	ab := &PrimsEdge{
		u:      a.id,
		v:      b.id,
		weight: 2,
	}
	ac := &PrimsEdge{
		u:      a.id,
		v:      c.id,
		weight: 3,
	}
	ad := &PrimsEdge{
		u:      a.id,
		v:      d.id,
		weight: 3,
	}

	bc := &PrimsEdge{
		u:      b.id,
		v:      c.id,
		weight: 4,
	}
	be := &PrimsEdge{
		u:      b.id,
		v:      e.id,
		weight: 3,
	}

	cd := &PrimsEdge{
		u:      c.id,
		v:      d.id,
		weight: 5,
	}
	cf := &PrimsEdge{
		u:      c.id,
		v:      f.id,
		weight: 6,
	}
	ce := &PrimsEdge{
		u:      c.id,
		v:      e.id,
		weight: 1,
	}

	df := &PrimsEdge{
		u:      d.id,
		v:      f.id,
		weight: 7,
	}

	ef := &PrimsEdge{
		u:      e.id,
		v:      f.id,
		weight: 8,
	}

	fg := &PrimsEdge{
		u:      f.id,
		v:      g.id,
		weight: 9,
	}

	a.connections = heap.HeapInsert[PrimsEdge, string](a.connections, ab, primsEdgeLt)
	a.connections = heap.HeapInsert[PrimsEdge, string](a.connections, ac, primsEdgeLt)
	a.connections = heap.HeapInsert[PrimsEdge, string](a.connections, ad, primsEdgeLt)
	b.connections = heap.HeapInsert[PrimsEdge, string](b.connections, bc, primsEdgeLt)
	b.connections = heap.HeapInsert[PrimsEdge, string](b.connections, be, primsEdgeLt)
	c.connections = heap.HeapInsert[PrimsEdge, string](c.connections, cd, primsEdgeLt)
	c.connections = heap.HeapInsert[PrimsEdge, string](c.connections, ce, primsEdgeLt)
	c.connections = heap.HeapInsert[PrimsEdge, string](c.connections, cf, primsEdgeLt)
	d.connections = heap.HeapInsert[PrimsEdge, string](d.connections, df, primsEdgeLt)
	e.connections = heap.HeapInsert[PrimsEdge, string](e.connections, ef, primsEdgeLt)
	f.connections = heap.HeapInsert[PrimsEdge, string](f.connections, fg, primsEdgeLt)
	actual, totalCost := PrimsMinSpanningTree([]*PrimsNode{a, b, c, d, e, f, g}) //Total cost should be 24

	if totalCost != 24 {
		t.Errorf("Actual total cost:%v, expected total cost:%v", totalCost, 24)
	}
	if len(actual) != 6 {
		t.Errorf("Actual # Edges:%v, Expected # Edges:%v", len(actual), 6)
	}

	expected := []*PrimsEdge{ab, ad, ac, ce, cf, fg}
	if !sets.SetEquality(actual, expected, edgeEq) {
		t.Errorf("Actual Edges:%v, Expected Edges:%v", actual, expected)

	}

}

func TestMinBottleneckSpanningTree(t *testing.T) {

	a := &PrimsNode{id: "a", connections: heap.New[PrimsEdge, string](extractor)}
	b := &PrimsNode{id: "b", connections: heap.New[PrimsEdge, string](extractor)}
	c := &PrimsNode{id: "c", connections: heap.New[PrimsEdge, string](extractor)}
	d := &PrimsNode{id: "d", connections: heap.New[PrimsEdge, string](extractor)}
	ab := &PrimsEdge{
		u:      a.id,
		v:      b.id,
		weight: 4,
	}
	ac := &PrimsEdge{
		u:      a.id,
		v:      c.id,
		weight: 2,
	}

	bd := &PrimsEdge{
		u:      b.id,
		v:      d.id,
		weight: 30,
	}

	cd := &PrimsEdge{
		u:      c.id,
		v:      d.id,
		weight: 31,
	}
	cb := &PrimsEdge{
		u:      c.id,
		v:      b.id,
		weight: 3,
	}

	a.connections = heap.HeapInsert[PrimsEdge, string](a.connections, ab, primsEdgeLt)
	a.connections = heap.HeapInsert[PrimsEdge, string](a.connections, ac, primsEdgeLt)
	b.connections = heap.HeapInsert[PrimsEdge, string](b.connections, bd, primsEdgeLt)
	c.connections = heap.HeapInsert[PrimsEdge, string](c.connections, cd, primsEdgeLt)
	c.connections = heap.HeapInsert[PrimsEdge, string](c.connections, cb, primsEdgeLt)
	actual, totalCost := PrimsMinSpanningTree([]*PrimsNode{a, b, c, d}) //Total cost should be 24

	if totalCost != 35 {
		t.Errorf("Actual total cost:%v, expected total cost:%v", totalCost, 35)
	}
	if len(actual) != 3 {
		t.Errorf("Actual # Edges:%v, Expected # Edges:%v", len(actual), 3)
	}

	expected := []*PrimsEdge{ac, bd, cb}
	if !sets.SetEquality(actual, expected, edgeEq) {
		t.Errorf("Actual Edges:%v, Expected Edges:%v", actual, expected)

	}
}

func TestAddSmallerEdge(t *testing.T) {

	a := &PrimsNode{id: "a", connections: heap.New[PrimsEdge, string](extractor)}
	b := &PrimsNode{id: "b", connections: heap.New[PrimsEdge, string](extractor)}
	c := &PrimsNode{id: "c", connections: heap.New[PrimsEdge, string](extractor)}
	d := &PrimsNode{id: "d", connections: heap.New[PrimsEdge, string](extractor)}
	ab := &PrimsEdge{
		u:      a.id,
		v:      b.id,
		weight: 2,
	}
	ad := &PrimsEdge{
		u:      a.id,
		v:      d.id,
		weight: 3,
	}

	bc := &PrimsEdge{
		u:      b.id,
		v:      c.id,
		weight: 4,
	}

	dc := &PrimsEdge{
		u:      d.id,
		v:      c.id,
		weight: 2,
	}

	a.connections = heap.HeapInsert[PrimsEdge, string](a.connections, ab, primsEdgeLt)
	a.connections = heap.HeapInsert[PrimsEdge, string](a.connections, ad, primsEdgeLt)
	b.connections = heap.HeapInsert[PrimsEdge, string](b.connections, bc, primsEdgeLt)
	d.connections = heap.HeapInsert[PrimsEdge, string](d.connections, dc, primsEdgeLt)

	//Changing the algorithm to return a map instead of an array allows me to O(1) lookup whether
	//or not the new edge ending in v would be in the minimum spanning tree. See the test below and
	//remove the commented-out section and comment out the additional edge db below. Then verify
	//this manually.  This is part (a) of question 10.
	//g := []*PrimsNode{a, b, c, d}
	//	actual := PrimsMinSpanningTreeReturningMap(g)
	//	fmt.Println(actual)
	db := &PrimsEdge{
		u:      d.id,
		v:      b.id,
		weight: 1,
	}
	d.connections = heap.HeapInsert[PrimsEdge, string](d.connections, db, primsEdgeLt)
	g := []*PrimsNode{a, b, c, d}

	actual := PrimsMinSpanningTreeReturningMap(g)
	fmt.Println(actual)
	if actual[db.v].weight < db.weight {
		fmt.Println("new edge would not be in tree")
	}
}
