package skiena_4

import (
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/sorting"
	"testing"
)

type Segment struct {
	l int
	r int
}

// Determine in Big O(N) the maximum number of Citations for authors.  It must be how group-by is implemented in a database.
/**
You are given a set S of n segments on the line, where Si ranges from Li to Ri.  Give an efficient
algorithm to select the fewest number of segments whose union completely covers the
interval 0 to m.

Answer:
Efficiency: O(n log n) + O(n) = O(n log n)
1. Sort the list by left point.
2. Iterate over the list.
3. Always choose the first segment and add it to result.
4. If next segment is not completely enclosed by previous one in result add it to result.
5. Repeat this process until you reach end of criteria.
6. return result
*/
func TestSmallestUnion(t *testing.T) {

	smallestUnion := func(allSegments []Segment, criteriaLimit int) []Segment {
		lt := func(l, r Segment) bool {
			if l.l < r.l {
				return true
			}
			return false
		}
		sorting.QuickSort(allSegments, lt)

		var result []Segment
		result = append(result, allSegments[0])
		for i := 0; i < len(allSegments); i++ {
			if i > 1 && allSegments[i].l > allSegments[i-1].r { //There is a gap
				return nil
			}
			prevR := result[(len(result) - 1)]
			//segment i is completely enclosed by previous one in result.
			if prevR.r >= allSegments[i].l {
				if prevR.r != allSegments[i].r && prevR.l != allSegments[i].l {
					result = append(result, allSegments[i])
				}
				if allSegments[i].r > criteriaLimit {
					break
				}
			}
		}
		return result
	}

	allSegments := []Segment{{2, 10}, {4, 12}, {0, 8},
		{8, 11}, {10, 13}, {16, 20}}
	actual := smallestUnion(allSegments, 17)
	if actual != nil {
		t.Errorf("Should have found gap")
	}
	allSegments = []Segment{{2, 10}, {0, 3}, {5, 12}}
	actual = smallestUnion(allSegments, 9)
	if actual == nil {
		t.Errorf("Should not have found gap")
	}

	/////////////////////////////////////////////////////////////
	allSegments = []Segment{{2, 10}, {0, 3}, {5, 12}}
	eq := func(l, r Segment) bool {
		if l.l == r.l && l.r == r.r {
			return true
		}
		return false
	}

	actual = smallestUnion(allSegments, 9)
	expected := []Segment{{0, 3}, {2, 10}}
	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("actual %v != expected %v", actual, expected)
	}

	/////////////////////////////////////////////////////////////
	allSegments = []Segment{{2, 8}, {0, 3}, {5, 12}}
	actual = smallestUnion(allSegments, 9)
	expected = []Segment{{0, 3}, {2, 8}, {5, 12}}
	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
	////////////////////////////////////
}
