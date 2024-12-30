The Problem:
Devise an algorithm for making a schedule that minimizes the total time it takes all contestants to complete a triathlon. The swimming part
is first for all contestants and is the only single-threaded part. Each contestant has an already known time to complete each leg. Here is the algorithm

Sort the contestants by the completion time of the first leg, the swimming leg, since that is the only leg that is single-threaded. 
And assign contestants to start in that order.  This will produce a schedule for all contestants that minimizes the completion time 
of the whole event. This algorithm is the one on page 167, the Minimize Lateness algorithm.