The Problem:
Given a graph of roads between villages in the mountains, produce a minimum spanning tree
consisting of the paths between villages that have the minimum maximum altitude.  The
goal is to identify the roads that need to be plowed by VDOT to keep all the towns
connected by one path.  And the higher the altitude, the more expensive it is to keep
the road plowed.

Prove that the following conjecture is true:
```
Conjecture 1: The minimum spanning tree of G, with respect to the edge weights A e, is a minimum-altitude connected subgraph.
Conjecture 2: A subgraph(V,E) is a minimum-altitude connected subgraph if and only if it contains the edges of the minimum spanning tree.
```

```
Cut Property: Assuming that all edge costs are distinct, let S be any subset of nodes that is neither empty nor equal to all
of v, and let edge e = (v, w) be the minimum-cost edge with one end in s and one end in V - S. Then every minimum spanning
tree contains the edge e.
```

Conjecture 2 is true because of the cut property.   And if conjecture 2 is true then conjecture 1 is also true,  since every
minimum spanning tree contains its own edges.
