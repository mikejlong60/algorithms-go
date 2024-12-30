package chapter4

import (
	"github.com/greymatter-io/golangz/heap"
	"github.com/greymatter-io/golangz/sets"
	"github.com/greymatter-io/golangz/sorting"
)

func kruskals(h heap.Heap[PrimsEdge, string], r map[string]*PrimsEdge, lt func(l, r *PrimsEdge) bool, expectedSize int) (heap.Heap[PrimsEdge, string], map[string]*PrimsEdge, func(l, r *PrimsEdge) bool, int) {
	if len(r) == expectedSize {
		return h, r, lt, expectedSize
	} else {
		a, _ := heap.FindMin(h)

		_, alreadySeen := r[a.v]
		if !alreadySeen {
			r[a.v] = a
		}
		h, _ = heap.HeapDelete(h, 0, lt)
		return kruskals(h, r, lt, expectedSize)
	}
}

func Kruskals(g []*PrimsEdge) []*PrimsEdge {

	// NOTE - this would have been less code if I had just used a sorted array instead of my heap.  But
	// maintaining familiarity with my heap is important because AAC uses it.
	toHeap := func(xs []*PrimsEdge, lt func(l, r *PrimsEdge) bool) heap.Heap[PrimsEdge, string] {
		exf := func(e *PrimsEdge) string {
			return e.v
		}
		h := heap.New[PrimsEdge, string](exf)

		for _, b := range xs {
			h = heap.HeapInsert(h, b, lt)
		}
		return h
	}

	lt := func(l, r *PrimsEdge) bool {
		if l.weight < r.weight {
			return true
		} else {
			return false
		}
	}

	toArray := func(xs map[string]*PrimsEdge) []*PrimsEdge {
		z := []*PrimsEdge{}
		for _, x := range xs {
			z = append(z, x)
		}
		return z
	}

	numberOfNodesMinus1 := func(xs []*PrimsEdge) int {
		var r = map[string]interface{}{}
		var c = 0
		for _, j := range xs {
			var there bool
			_, there = r[j.u]
			if !there {
				r[j.u] = nil
				c = c + 1
			}
			_, there = r[j.v]
			if !there {
				r[j.v] = nil
				c = c + 1
			}
		}
		return c - 1
	}
	_, r, _, _ := kruskals(toHeap(g, lt), map[string]*PrimsEdge{}, lt, numberOfNodesMinus1(g))

	return toArray(r)
}

func KruskalsUsingUnionFind(g []*PrimsEdge) []*PrimsEdge {
	ltpe := func(l, r *PrimsEdge) bool {
		if l.weight < r.weight {
			return true
		} else {
			return false
		}
	}

	sorting.QuickSort(g, ltpe)

	lt := func(l, r string) bool {
		if l < r {
			return true
		} else {
			return false
		}
	}

	eq := func(l, r string) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}

	makeSetOfNodeIds := func(g []*PrimsEdge) []string {
		r := []string{}
		for _, b := range g {
			r = append(r, b.u)
			r = append(r, b.v)
		}
		return sets.ToSet(r, lt, eq)
	}

	uf := MakeUnionFind(makeSetOfNodeIds(g))

	//Put UNodes into a map so you can look them up
	ufm := map[string]*UNode{}
	for _, b := range uf {
		ufm[b.Id] = b
	}

	//Iterate over g and decide whether or not to use the edge in the minimum spanning tree.
	r := []*PrimsEdge{}
	for _, b := range g { //g is sorted-by-weight array of edges
		vId := Find(ufm[b.v])
		if vId == b.v { //If b is not yet in a set, add it to a set and include it in the MST.
			r = append(r, b)
			Union(ufm[b.u], ufm[b.v])
		}
	}

	return r
}
