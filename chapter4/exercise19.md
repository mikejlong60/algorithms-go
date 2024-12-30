The Problem:
Given a graph of edges(all water pipes are of varying diameter) that represent drinking water pumping stations linked by
pipes all of whom draw water from the same intake pipe on the Potomac river, give a polynomial-time algorithm that maximizes the
flow rate between any two pumping stations.  The

The resulting tree will show the optimal path from a pumping station to every other pumping station.

Solution:
My Prims implementation will calculate this with a slight variation; that being the length value of the edge(interpreted here as the flow rate)
should be it's reciprocal since we are trying to maximize the total volume of water that flows through the graph.

Algorithm efficiency O(m og n) where m is the number of edges and n is the number iof nodes.

This is the Prims algorithm to compute the minimum spanning tree.