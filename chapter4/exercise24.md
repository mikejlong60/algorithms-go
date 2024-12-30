Problem:
Starting with a Complete Balanced Binary Tree with a weight on each edge.  Think of it as a circuit with expensive wires connecting
each node. The goal is to use the minimum amount of wire to connect all the nodes. And the path from the root to every node must be the
same length because we want a signal from the root to reach each leaf node at the same time.  This is known as the
zero-skew problem.

// Algorithm

// 1. Build a balanced binary tree(even number of leaf nodes) using Huffman with equal probabilities for each letter.  Assign a ramdom integer
// that serves as a weight to each edge in tree as you build it.
// 2. Derive the list of leaf nodes with their accumulated weights using a variation of Depth-First Search from chapter 3.
// 3. Sort the list of leaf nodes by highest weight sum descending.
// 4. Map over xs and increase the edge above each leaf node so that its sum equals the first leaf node(the most expensive).


Problem. This approach is flawed because I am not guaranteeing the shortest amount of wire for the whole tree, only that the distance
from root to each leaf is the same. In order to correct this I need an algorithm that looks up from each leaf and finds the most general edge
to change such that all leafs under that have the same total length, and that set of leaves has the weight of the highest edge.
That is too hard for me at the moment.
