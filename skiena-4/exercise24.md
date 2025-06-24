What is the worst-case running time of merge sort if you divide the array to be sorted into 
thirds and merge all three sub-arrays at each step.

Here is how that algorithm compares with regular merge sort that splits the array in half until it
reaches  two elements, then merges the two arrays in the layer above.  Here is a comparison

merging an array of 2^5 integers
Mergesort-divide in half                                   
1.Split array in half.  You have 5 layers.
2. At 5 bottom layer have 16 2-element arrays. Worst case 32 comparisons total
3. Do all permutations to order the two numbers(2!) and return the order array to above layer.
3. Now there are 8 arrays after the merge operation and return the ordered array to above layer. That merge is O(N)
4. Now there are 4 arrays. Return those 4 arrays to above layer. That array merge is O(n).
5. Now there are 2 arrays  That array merge is O(n).
6. Return the sorted array.
Cost of this sort is O(log^3 of n + (n/3 * n log^3 of n factorial))
   Same as standard merge sort — the base of the logarithm changes (from 2 to 3), but asymptotically, it’s still
   log
   ⁡
   𝑛
   logn.

Mergesort-divide in thirds                                   
1.Split array in thirds.  You have 3 layers.
2. At 3 bottom layer you have 12 3-element arrays.
3. Do 3! or 6 comparisons instead of 1 to order the three numbers and return the order array to above layer.  Worst case 72 comparisons.
3. Now there are 6 arrays after the merge operation and return the ordered array to above layer. That merge is O(N)
4. Now there are 3 arrays. Return those 4 arrays to above layer. That array merge is O(n).
5. Return the sorted array. 

Cost of merge is O(log^3 n)
Total time: proportional to 𝑁