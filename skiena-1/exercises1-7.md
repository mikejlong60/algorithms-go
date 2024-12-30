# Finding Counterexamples
  1. Show that a + b can be less than min(a, b).
    * Answer: -2 + -10 = -12.  -12 is less than -10
  2. Show that a * b is less than min(a, b).
    * Answer: 2 * -10 is -20. -20 is less than  -10
  3. Design a road network with two points a and b such that the fastest route is not the shortest one.
    * Answer: Route 1 from Winchester to Norfolk is 150 miles. Route 95 from Winchester to Norfolk is 175 miles. Route 1 has a speed limit of 50 mph
       Route 95 has a speed limit of 70 mph. It takes 3 hours to go route 1, 150/50  = 3 hours.  Route 95 takes 175/70 = 2.5 hours.
  4. Design a road network with two points such that the shortest route between is not the one with the least turns.
    * Answer: Pick a point for A and go north and west for 1 inch at a 45 degree angle.  And point B is 1 inch east from that point where A's first
       segment

          |
            |       |
              |        |
                a        b

               This following is a longer distance than above even though is has 3 turns instead of 2.


                   |   |
               a |   |   |
                           b

  5. Given a set of integers S = {a, b, c, d} and a target number T, find a subset of S that adds up to exactly T.
      Example set S = {1,2,5,9,10}.   There exists a subset of S where T = 22 but not T = 23
      Find counter examples for the following algorithms:
    * Answer:
      * Put the elements in the knapsack in left to right order. There is no T = 22.
      * Put the elements in the knapsack from smallest to largest. There is no T = 22.
      * Put the elements in the knapsack from largest to smallest. There is no T = 22.
  6. Set cover problem. Given a set S of subsets S1, ... Sn, find the smallest subset of subsets such that the elements
     in that set of subsets includes every element in the original set of subsets.  For example the subsets
     S1 = {1,3,5}, S2 = {2,4}, S3 = {1,4}, s4 = {2,5}.  In this example the sets S1 and S2 would be the smallest subset
     of subsets.

     Now give a counterexample which refutes the following algorithm:
     Select the largest subset, and then delete all its elements from the universal set.  And then add the subset containing
     the largest number of uncovered elements until all are covered.
    * Answer: This algorithm works for the example above.  But this algorithm does not consider
      the possibility of ties for the highest number of elements in the subsets.
        S1 = {1,4}, S2 = {1,4,5}, S3 = {1,3}, S4 = {1,2,5,6}.  This example would fail because after first choosing S4,
        the algorithm might choose any of S1, S2 or S3 which all tie in covering 1 element in the remaining set. If S2 is chosen
        instead of S1 the answer would be S2, S3, and S4 instead of S1, S3, and S4.  The inclusion of an additional rule
        which chooses the set with the smallest number of covering elements in the event of a tie would correct this error.
  7. Maximum clique problem - Find the largest subset of vertices such that there is an edge between them all.
     Find a counter example for the following algorithm:
        Sort the vertices by degree descending. Iterate over those vertices and add the vertex to the clique if it is a
        neighbor of all other vertices in the clique.

        a --- b     e --- f
         |   |       |   |
           c --------- d

           That algorithm would produce an answer of c --- d being the maximum clique when the correct answer would be a --- b --- c or e --- f --- d.