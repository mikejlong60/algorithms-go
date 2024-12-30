package chapter3

import "github.com/greymatter-io/golangz/propcheck"

type NodeForTopoOrdering struct {
	Id                  int
	OutgoingConnections map[int]*NodeForTopoOrdering
	IncomingConnections map[int]*NodeForTopoOrdering
}

func Topo(m map[int]*NodeForTopoOrdering, accum []*NodeForTopoOrdering) (map[int]*NodeForTopoOrdering, []*NodeForTopoOrdering) {
	if len(m) == 0 {
		return m, accum
	} else {
		//Find next node `n` in map `m` with no incoming edges
		var n *NodeForTopoOrdering
		for _, j := range m {
			if len(j.IncomingConnections) == 0 {
				n = j
				//Iterate over each `n`'s outgoing connections and remove `n` from `p`'s list of incoming connections
				for _, p := range n.OutgoingConnections {
					delete(p.IncomingConnections, n.Id)
				}
				break
			}
		}
		//Append `n` to accum
		accum = append(accum, n)
		//Remove `n` from `m`
		delete(m, n.Id)
		return Topo(m, accum)
	}
}

func MakeConnectedComponentsAsNodeForTopoOrdering(a propcheck.Pair[map[int]*Node, int]) map[int]*NodeForTopoOrdering {
	graph := make(map[int]NodeForTopoOrdering, len(a.A))
	for _, xs := range a.A { //Convert initial list of nodes to the type from which you can make a topological ordering.
		ie := make(map[int]*NodeForTopoOrdering)
		oe := make(map[int]*NodeForTopoOrdering)
		graph[xs.Id] = NodeForTopoOrdering{Id: xs.Id, IncomingConnections: ie, OutgoingConnections: oe}
	}

	cc := GenerateConnectedComponents(a)
	var nodes = make(map[int]*NodeForTopoOrdering)
	for _, xs := range cc.A {
		if len(xs) > 1 && len(xs) < 4 { //Transform first connected component graph that is larger than one node into its equivalent NodeForTopoOrdering map for computing Topo ordering
			for _, ys := range xs {
				n := graph[ys.U]
				oe := graph[ys.V]
				oe.IncomingConnections[n.Id] = &n
				n.OutgoingConnections[ys.V] = &oe
				if ys.U != ys.V { //Don't add top-level edge that points to itself
					n.IncomingConnections[ys.U] = &n
				}
				nodes[n.Id] = &n
			}
			break
		}
	}
	return nodes
}
