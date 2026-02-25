Question: Prove that a priority queue can not have  INSERT, MAX, and EXTRACT MAX
operations in O(1) worst case as a fool claims. And the reason that this
fool's claim is false is due to the O(n log n) lower bound for sorting.

Answer: MAX is possible in O(1) because it does not require HEAPIFY UP or 
HEAPIFY DOWN for the priority queue which is a O(log n) operation.  
But the EXTRACT MAX and INSERT are both O(log n) operations.  Therefore 
the O(1) claim is incorrect.



//////////////////
Bottom line: You’re wrong that insertion “always” takes
𝑂(log ⁡ 𝑛) O(logn). It’s 𝑂(log ⁡ 𝑛)O(logn) for heaps, but priority queues can have 𝑂(1)
O(1) insertion (amortized), and the real theorem is: you can’t have INSERT + MAX + EXTRACT-MAX all 𝑂(1)
O(1) worst-case in the comparison model, because that would give you an 𝑂(𝑛)O(n) comparison sort.