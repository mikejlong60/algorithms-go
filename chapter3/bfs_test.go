package chapter3

import (
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/propcheck"
	"testing"
	"time"
)

func TestHasCycle(t *testing.T) {
	n0 := Node{
		Id:          0,
		Connections: nil,
	}
	n1 := Node{
		Id:          1,
		Connections: nil,
	}
	n2 := Node{
		Id:          2,
		Connections: nil,
	}
	n3 := Node{
		Id:          03,
		Connections: nil,
	}
	n4 := Node{
		Id:          4,
		Connections: nil,
	}
	n5 := Node{
		Id:          5,
		Connections: nil,
	}
	n6 := Node{
		Id:          6,
		Connections: nil,
	}
	n7 := Node{
		Id:          7,
		Connections: nil,
	}
	n0.Connections = []*Node{&n1, &n2}
	n1.Connections = []*Node{&n4}
	n2.Connections = []*Node{&n3}
	n3.Connections = []*Node{&n5, &n6, &n7}
	n7.Connections = []*Node{&n7, &n2}
	n5.Connections = []*Node{&n7}
	n6.Connections = []*Node{&n4}
	graph := make(map[int]*Node, 7) //First field of pair is the layer the node is in, -1 means it's never been seen before and is thus not in any layer
	graph[0] = &n0
	graph[1] = &n1
	graph[2] = &n2
	graph[3] = &n3
	graph[4] = &n4
	graph[5] = &n5
	graph[6] = &n6
	graph[7] = &n7
	tree, hasCycle, numNodes := BFSearch(graph, 0)
	expected := [][]Edge{{{-1, 0}}, {{0, 2}, {0, 1}}, {{1, 4}, {2, 3}}, {{3, 5}, {3, 6}, {3, 7}}}
	if !arrays.ArrayEquality(tree, expected, TreeEquality) {
		t.Errorf("Actual:%v Expected:%v", tree, expected)
	}
	if !hasCycle {
		t.Errorf("tree should have had a cycle")
	}
	if len(graph) != numNodes {
		t.Errorf("Tree had %v nodes and expected it to have %v nodes", numNodes, len(graph))
	}
	hasCycle, N_1EdgesAndConnected, err := Rule3_2(graph, 0)
	if N_1EdgesAndConnected && hasCycle {
		t.Errorf("Rule 3.2 failure:%v", err)
	}
}

func TestRule3_2(t *testing.T) {
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}

	prop := propcheck.ForAll(UndirectedGraphGen(1, 100),
		"Generate a random graph and do a Tree search starting from some root.",
		func(graph propcheck.Pair[map[int]*Node, int]) propcheck.Pair[propcheck.Pair[bool, bool], string] {
			hasCycle, N_1EdgesAndConnected, err := Rule3_2(graph.A, graph.B)
			a := propcheck.Pair[bool, bool]{hasCycle, N_1EdgesAndConnected}
			return propcheck.Pair[propcheck.Pair[bool, bool], string]{a, err}
		},
		func(p propcheck.Pair[propcheck.Pair[bool, bool], string]) (bool, error) {
			var errors error
			if !p.A.B {
				t.Errorf("Rule 3 failure:%v", p.B)
			}
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{50, rng})
	propcheck.ExpectSuccess[propcheck.Pair[map[int]*Node, int]](t, result)
}
