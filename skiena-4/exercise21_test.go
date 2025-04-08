package skiena_4

import (
	"math/rand"
	"testing"
)

//  Find the median of a set of numbers using some sort of partitioning technique like Quicksort.
/** Algorithm
1. Pick a pivot point in the array. Can be anything
2. Iterate through the array, swapping the pivot value until you have placed it in the correct index of the array.
   You can have more than 1 value, the array is not a set. This is the array for this example: {5, 2, 5, 1, 5, 3, 9, 5}.
   Your initial pivot is index 0, the value 5.  But the choice of pivot point does not matter.
3. Once you have done that you end up with three sub-arrays: the one which contains the pivot value, in this case all the 5s,
   the one to the left which contains everything less than the pivot value, and the one to the right which contains
   everything greater than the pivot value.
*/

func partition(xs []int, left, right int) int {
	pivotIndex := rand.Intn(right-left+1) + left
	//NOTE - This a special double assignment statement.
	//It evaluates 2 right side variables first and then assigns to left
	//side without the need for a temporary variable like I would need in C.
	xs[pivotIndex], xs[right] = xs[right], xs[pivotIndex]
	pivot := xs[right]
	storeIndex := left

	for i := left; i < right; i++ {
		//This does the swap
		if xs[i] < pivot {
			//Double assignment statement.
			xs[storeIndex], xs[i] = xs[i], xs[storeIndex]
			storeIndex++
		}
	}
	//Double assignment statement.
	xs[right], xs[storeIndex] = xs[storeIndex], xs[right]
	return storeIndex
}

func quickSelect(xs []int, left, right, k int) int {
	if left == right {
		return xs[left]
	}

	pivotIndex := partition(xs, left, right)
	if k == pivotIndex {
		return xs[k]
	} else if k < pivotIndex {
		return quickSelect(xs, left, pivotIndex-1, k)
	} else {
		return quickSelect(xs, pivotIndex+1, right, k)
	}
}

func TestQuickSelect(t *testing.T) {

	var actual int
	xs := []int{7, 2, 1, 6, 8, 5, 3, 4}

	n := len(xs)
	k := n / 2
	if n%2 == 1 {
		//odd number of elements
		actual = quickSelect(xs, 0, n-1, k)
	} else {
		//even number of elements
		left := quickSelect(xs, 0, n-1, k-1)
		right := quickSelect(xs, 0, n-1, k)
		actual = (left + right) / 2
	}

	expected := 4
	if expected != actual {
		t.Errorf("Expected:[%v], Actual:[%v]", expected, actual)
	}

}
