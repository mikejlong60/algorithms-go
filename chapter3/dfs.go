package chapter3

import (
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/greymatter-io/golangz/sets"
)

//Depth-First search
// A recursive algorithm for depth-first search.
//Params:
//  U - *Node the current node that gets explored by the algorithm
//  seen - seen map[int]*Node - the accumulated map of Nodes that the algorithm has seen thus far
//  tree- an array of Edges reflecting the current dfs tree to this point
//Returns:
//  U - *Node the current node that gets explored by the algorithm
//  seen - seen map[int]*Node - the accumulated map of Nodes that the algorithm has seen thus far
//  tree- an array of Edges reflecting the current dfs tree to this point

func DFSearch(u *Node, seen map[int]*Node, tree []Edge) (*Node, map[int]*Node, []Edge) {
	seen[u.Id] = u
	for _, connectedNode := range u.Connections {
		_, explored := seen[connectedNode.Id]
		if !explored {
			tree = append(tree, Edge{u.Id, connectedNode.Id})
			_, seen, tree = DFSearch(connectedNode, seen, tree)
		}
	}
	return u, seen, tree
}

func GenerateConnectedComponents(graph propcheck.Pair[map[int]*Node, int]) propcheck.Pair[[][]Edge, []int] {
	lt := func(l, r int) bool {
		if l < r {
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

	var tree []Edge
	makeAllUnseenNodes := func(nodeMap map[int]*Node) []int {
		var r []int
		for i, _ := range nodeMap {
			r = append(r, i)
		}
		return r
	}
	var unseenNodes = makeAllUnseenNodes(graph.A)

	toNodeIdSet := func(tree []Edge) []int {
		var r []int
		for _, node := range tree {
			r = append(r, node.V)
			r = append(r, node.U)
		}
		r = sets.ToSet(r, lt, eq)
		return r
	}

	var allConnectedComponents [][]Edge
	for len(unseenNodes) > 0 {
		_, _, dfTree := DFSearch(graph.A[unseenNodes[0]], make(map[int]*Node), tree) //Build a tree from the first node of the unseen nodes.
		if len(dfTree) == 0 {                                                        //If tree is empty the node has no connections but is still connected to itself
			dfTree = []Edge{{unseenNodes[0], unseenNodes[0]}}
		}
		allConnectedComponents = append(allConnectedComponents, dfTree)
		dfNodes := toNodeIdSet(dfTree)
		unseenNodes = sets.SetMinus(unseenNodes, dfNodes, eq)
	}
	allNodes := makeAllUnseenNodes(graph.A)
	return propcheck.Pair[[][]Edge, []int]{allConnectedComponents, allNodes}
}

func MakeConnectionComponents(graph map[int]*Node) propcheck.Pair[[][]Edge, []int] {

	lt := func(l, r int) bool {
		if l < r {
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

	var tree []Edge
	makeAllUnseenNodes := func(nodeMap map[int]*Node) []int {
		var r []int
		for i, _ := range nodeMap {
			r = append(r, i)
		}
		return r
	}
	var unseenNodes = makeAllUnseenNodes(graph)

	toNodeIdSet := func(tree []Edge) []int {
		var r []int
		for _, node := range tree {
			r = append(r, node.V)
			r = append(r, node.U)
		}
		r = sets.ToSet(r, lt, eq)
		return r
	}

	var allConnectedComponents [][]Edge
	for len(unseenNodes) > 0 {
		_, _, dfTree := DFSearch(graph[unseenNodes[0]], make(map[int]*Node), tree) //Build a tree from the first node of the unseen nodes.
		if len(dfTree) == 0 {                                                      //If tree is empty the node has no connections but is still connected to itself
			dfTree = []Edge{{unseenNodes[0], unseenNodes[0]}}
		}
		allConnectedComponents = append(allConnectedComponents, dfTree)
		dfNodes := toNodeIdSet(dfTree)
		unseenNodes = sets.SetMinus(unseenNodes, dfNodes, eq)
	}
	allNodes := makeAllUnseenNodes(graph)
	return propcheck.Pair[[][]Edge, []int]{allConnectedComponents, allNodes}
}
