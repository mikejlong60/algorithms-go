package skiena_4

// Show that you can sort an array of k distinct integers in O(n log k) steps, better than O(n log n).  Think of sorting
// an array of 10,000 ones and zeros.

//This can be shown to be true using a heap. Since any value can be inserted into a heap in log n steps where n is the
//size of the heap.  In the case of the examole above where k is two, the algorithm is as follows:

/**
heap y
for i, _ := range xs { // O(log k) * n total steps.
	incrementI(y, i)
}

Now just make a new array by popping off the heap until it is empty and writing number of increment elements for k.

I think this must be how heap sort is implemented.

*/
