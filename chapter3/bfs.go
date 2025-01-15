package chapter3

import (
	"fmt"
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
)

type Node struct {
	Id          int
	Connections []*Node
}

type Edge struct {
	U int //the Id of the beginning node of the edge
	V int //the Id of the ending node of the edge
}

type NodeLayerTuple struct {
	Id               int //The node Id
	DistanceFromRoot int //Only used by the Djkstr4a algorithm in chapter 4
	Layer            int // Zero indexed array index indicating the layer(array index) in the Tree array the node lives
}

// Breadth-First search with cycle detection
// Params:
//
//	graph a hashmap of all the nodes in te graph. Facilitates n log n lookup
//	rootId the Node Id of the root node, the one at the top of the mobile from which all the other nodes hang
//
// Returns:
//
//	Tree  - the search tree represented as an array of layers, each layer constisting of an array of Edges(U, V)
//	bool - whether or not the resulting search tree contained a cycle. A cycle is a relationship between two nodes that is farther than one layer apart.
//	int - the number of nodes in the Tree
func BFSearch(graph map[int]*Node, rootId int) ([][]Edge, bool, int) {
	hasCycle := func(nodeId int, currentLayer int, layers map[int]NodeLayerTuple) bool {
		l := layers[nodeId]
		if currentLayer-2 >= l.Layer { //there is a cycle
			return true
		} else {
			return false
		}
	}

	var tree = [][]Edge{}
	l0 := []Edge{{U: -1, V: rootId}}

	//A lookup map so you can look up whether or not a Node has been seen and if so what layer it is in.
	layersLookup := make(map[int]NodeLayerTuple, len(graph))
	layersLookup[rootId] = NodeLayerTuple{
		Id:    rootId,
		Layer: 0,
	}
	tree = append(tree, l0)

	var graphHasACycle = false
	var i = 0 //current layer you are finding edges for
	for {
		var pendingLayer []Edge
		for _, k := range tree[i] {
			node, _ := graph[k.V]
			for _, m := range node.Connections {
				//Lookup tail(V) of every edge in the layer to see if it has been seen before. If not add it to pending layer.
				_, alreadySeen := layersLookup[m.Id]
				if !alreadySeen {
					pendingLayer = append(pendingLayer, Edge{U: k.V, V: m.Id})
					layersLookup[m.Id] = NodeLayerTuple{Id: m.Id, Layer: i + 1}
				} else {                 //Don't add it since we already know about this Node. But DO see if its a cycle.
					if !graphHasACycle { //Can only set this value to true one time
						graphHasACycle = hasCycle(m.Id, i, layersLookup)
					}
				}
			}
		}
		if len(pendingLayer) > 0 {
			tree = append(tree, pendingLayer)
			i = i + 1
		} else {
			break
		}
	}
	return tree, graphHasACycle, len(layersLookup)
}

func TreeEquality(a, b []Edge) bool {
	edgeEq := func(a, b Edge) bool {
		if a.U == b.U && a.V == b.V {
			return true
		} else {
			return false
		}
	}
	if sets.SetEquality(a, b, edgeEq) {
		return true
	} else {
		return false
	}
}

func Rule3_2(graph map[int]*Node, rootNode int) (bool, bool, string) {
	bfsTree, hasCycle, numNodes := BFSearch(graph, rootNode)
	numEdgesInTree := func(tree [][]Edge) int {
		var edges int
		for _, node := range tree {
			edges = edges + len(node)
		}
		return edges - 1
	}

	numEdges := numEdgesInTree(bfsTree)
	isConnected := true
	hasN_1Edges := numNodes-1 == numEdges

	//hasCycle is based upon the original graph. The resulting bfsTree has no cycles
	return !hasCycle, isConnected && hasN_1Edges, fmt.Sprintf("Has Cycle:%v, isConnected: %v, has n-1 edges:%v\n:", hasCycle, isConnected, hasN_1Edges)
}

func UndirectedGraphGen(lower, upperExc int) func(propcheck.SimpleRNG) (propcheck.Pair[map[int]*Node, int], propcheck.SimpleRNG) {
	return func(rng propcheck.SimpleRNG) (propcheck.Pair[map[int]*Node, int], propcheck.SimpleRNG) {
		nodeEq := func(l, r *Node) bool {
			if l.Id == r.Id {
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

		lt := func(l, r int) bool {
			if l < r {
				return true
			} else {
				return false
			}
		}

		nodeIds, rng2 := sets.ChooseSet(lower, upperExc, propcheck.ChooseInt(0, 1000000), lt, eq)(rng)
		graph := make(map[int]*Node, len(nodeIds))
		for _, j := range nodeIds {
			graph[j] = &Node{Id: j}
		}

		var rng3 = rng2
		var connectionIds []int
		for _, node := range graph {
			var connections []*Node
			connectedNodeSize := len(nodeIds)
			connectionIds, rng3 = sets.ChooseSet(0, int(connectedNodeSize), propcheck.ChooseInt(0, int(connectedNodeSize)), lt, eq)(rng3)
			for _, connectedNodeId := range connectionIds {
				if node.Id != graph[nodeIds[connectedNodeId]].Id {
					connections = append(connections, graph[nodeIds[connectedNodeId]])
				}
			}
			node.Connections = connections
		}
		for _, node := range graph {
			////Now make sure every node's connections array is connected to the node to which it points from the other node's perspective
			for _, conn := range node.Connections {
				if !arrays.Contains(conn.Connections, node, nodeEq) {
					conn.Connections = append(conn.Connections, node)
				}
			}
		}
		root, rng4 := propcheck.ChooseInt(0, len(graph))(rng3)
		return propcheck.Pair[map[int]*Node, int]{graph, nodeIds[root]}, rng4
	}
}

func EvenNumberOfNodesGen(lower, upperExc int) func(propcheck.SimpleRNG) (map[int]*Node, propcheck.SimpleRNG) {
	nodeEq := func(l, r *Node) bool {
		if l.Id == r.Id {
			return true
		} else {
			return false
		}
	}
	nodeLt := func(l, r *Node) bool {
		if l.Id < r.Id {
			return true
		} else {
			return false
		}
	}

	return func(rng propcheck.SimpleRNG) (map[int]*Node, propcheck.SimpleRNG) {

		eq := func(l, r int) bool {
			if l == r {
				return true
			} else {
				return false
			}
		}

		lt := func(l, r int) bool {
			if l < r {
				return true
			} else {
				return false
			}
		}

		nodeIds, rng2 := sets.ChooseSet(lower, upperExc, propcheck.ChooseInt(0, 1000000), lt, eq)(rng)
		graph := make(map[int]*Node, len(nodeIds))
		for _, j := range nodeIds {
			graph[j] = &Node{Id: j}
		}
		if len(graph) > 0 && len(graph)%2 != 0 {
			delete(graph, nodeIds[0])
		}

		var nodes []*Node
		for _, j := range graph {
			nodes = append(nodes, j)
		}

		//Now connect each node to at least half of the other nodes
		var rng3 = rng2
		var connections []int
		for _, j := range graph {
			connections, rng3 = sets.ChooseSet(len(nodes)/2, len(nodes), propcheck.ChooseInt(0, len(nodes)), lt, eq)(rng3)
			for _, l := range connections {
				graph[j.Id].Connections = append(graph[j.Id].Connections, nodes[l])
			}
			var connectionsSet []*Node
			if len(graph[j.Id].Connections) < len(nodes)/2 { //Add elements until you reach half connections because your previous set operation did not get enough.
				connectionsSet = sets.ToSet(graph[j.Id].Connections, nodeLt, nodeEq)
				for _, y := range graph {
					if len(connectionsSet) >= len(nodes)/2 {
						break
					} else {
						connectionsSet = sets.SetUnion(connectionsSet, nodeLt, eq, []*Node{y})
					}
				}
				graph[j.Id].Connections = connectionsSet
			}
		}
		return graph, rng3
	}
}
