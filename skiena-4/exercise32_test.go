package skiena_4

import (
	"testing"

	"github.com/greymatter-io/golangz/arrays"
)

//Implement wiggle sort.  The definition of wiggle sort is make an unsorted array of integers
//a ordered such that a[0] < a[1] and a[2] < a[1] and a[2] < a[3]

// Do it in O(n) in space O(1)
//Possible answers for a == [3,1,4,2,6,5] are [1,3,2,5,4,6]

//1,3,2,4,6,5

func wiggleSort(xs []int) []int {
	for i := 0; i < len(xs); i++ {
		if i < len(xs)-1 && xs[i] > xs[i+1] {
			xs[i], xs[i+1] = xs[i+1], xs[i]
		}
	}

	return xs
}

func TestWiggleSort(t *testing.T) {
	xs := []int{3, 1, 4, 2, 6, 5}

	actual := wiggleSort(xs)
	expected := []int{1, 3, 2, 5, 4, 6}
	eq := func(l, r int) bool {
		if l == r {
			return true
		}
		return false
	}
	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}
