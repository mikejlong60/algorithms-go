package skiena_4

import (
	"testing"

	"github.com/greymatter-io/golangz/arrays"
)

//Implement wiggle sort.  The definition of wiggle sort is make an unsorted array of integers
//ordered such that a[0] < a[1] and a[2] < a[1] and a[2] < a[3]

// Do it in O(n) in space O(1)
//Possible answers for a == [3,1,4,2,6,5] are [1,3,2,5,4,6]

//1,3,2,4,6,5

func wiggleSort(xs []int) []int {
	for i := 0; i < len(xs); i++ {
		if i < len(xs)-2 && xs[i] > xs[i+1] && xs[i] < xs[i+2] { //1,3,2,4 case
			xs[i], xs[i+1] = xs[i+1], xs[i]
			continue
		}
		if i < len(xs)-3 && xs[i] > xs[i+1] && xs[i] > xs[i+2] { //4,2,6,5 case
			//swap 2[i+1],4[i] then you have 2,4,6,5, then swap 4[i+1] and 5[i+3]
			xs[i+1], xs[i] = xs[i], xs[i+1]
			xs[i+1], xs[i+3] = xs[i+3], xs[i+1]
			continue
		}
		if i < len(xs)-1 && xs[i] > xs[i+1] { // now you can fall through
			xs[i], xs[i+1] = xs[i+1], xs[i]
			continue
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
