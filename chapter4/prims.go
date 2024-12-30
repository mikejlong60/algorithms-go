package chapter4

import (
	"fmt"
	"github.com/greymatter-io/golangz/heap"
	"math"
)

type PrimsEdge struct {
	u      string
	v      string
	weight float32
}

type PrimsNode struct {
	id          string
	connections heap.Heap[PrimsEdge, string]
}

func (w PrimsNode) String() string {
	return fmt.Sprintf("PrimsNode{id:%v, connections:%v}", w.id, w.connections)
}
func (w PrimsEdge) String() string {
	return fmt.Sprintf("PrimsEdge{u:%v, v:%v, weight:%v}", w.u, w.v, w.weight)
}

func primsEdgeLt(l, r *PrimsEdge) bool {
	if l.weight < r.weight {
		return true
	} else {
		return false
	}
}

func extractor(edge *PrimsEdge) string {
	return edge.v
}

func primsMinSpanningTree(xs []*PrimsNode, xxs []*PrimsEdge) ([]*PrimsNode, []*PrimsEdge) {
	//1. Is minEdge == nil return
	//2. Find minimum-cost connected edge(minEdge) among all edges in array of PrimsNodes.
	//3. Delete all edges that point to minEdge.v in original xs array.
	//4. Add that minEdge to PrimsEdge result (xxs) array
	//4. call primsMinSpanningTree again with updated xs and xxs

	deleteAllEdgesPointingToV := func(xs []*PrimsNode, v *PrimsEdge) []*PrimsNode {
		for _, y := range xs {
			p := heap.FindPosition(y.connections, v.v)
			if p > -1 {
				y.connections, _ = heap.HeapDelete(y.connections, p, primsEdgeLt)
			}
		}
		return xs
	}

	minEdge := func(xxs []*PrimsNode) *PrimsEdge {
		var lowestEdge = &PrimsEdge{
			u:      "",
			v:      "",
			weight: math.MaxInt,
		}
		for _, y := range xxs {
			a, err := heap.FindMin(y.connections)
			if err == nil && a.weight < lowestEdge.weight {
				lowestEdge = a
			}
		}
		if lowestEdge.weight == math.MaxInt {
			return nil
		} else {
			return lowestEdge
		}
	}

	e := minEdge(xs)
	if e == nil {
		return xs, xxs
	}
	xs = deleteAllEdgesPointingToV(xs, e)
	xxs = append(xxs, e)
	return primsMinSpanningTree(xs, xxs)
}

// Prims algorithm assumes that there are no cycles.
// It starts with the minimum edge but you could start with any node that has an outgoing edge.
// O(m log n) where  n is the number of nodes and m is the number of edges given that I have used my heap.
func PrimsMinSpanningTree(xs []*PrimsNode) ([]*PrimsEdge, float32) {
	totalCost := func(xs []*PrimsEdge) float32 {
		var r float32
		for _, b := range xs {
			r = b.weight + r
		}
		return r
	}
	_, r := primsMinSpanningTree(xs, []*PrimsEdge{})
	return r, totalCost(r)
}
