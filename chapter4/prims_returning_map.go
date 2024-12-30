package chapter4

func PrimsMinSpanningTreeReturningMap(xs []*PrimsNode) map[string]*PrimsEdge {
	_, r := primsMinSpanningTree(xs, []*PrimsEdge{})

	rr := make(map[string]*PrimsEdge)
	for _, b := range r {
		rr[b.v] = b
	}
	return rr
}
