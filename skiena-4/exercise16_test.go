package skiena_4

import (
	"github.com/greymatter-io/golangz/sorting"
	"testing"
)

type Segment struct {
	l int
	r int
}

// Determine in Big O(N) the maximum number of Citations for authors.  It must be how group-by is implemented in a database.
/**
You are given a set S of n segments on the line, where Si ranges from Li to Ri.  Give an effecient
algorithm to select the fewest number of segments whose union completely covers the
interval 0 to m.

Answer:
1. Sort the list by left point.
2. Iterate over the list.
3. Always choose the first segment and add it to result.
4. Start looping at 2nd element in array.
4. If next segment is not completely enclosed by previous one in result add it to result.
5. Repeat this process until you reach end of criteria.
6. return result
*/
func TestFewestSegments(t *testing.T) {

	lt := func(l, r Segment) bool {
		if l.l < r.l {
			return true
		} else {
			return false
		}
	}
	isThereAGap := func(allSegments []Segment, criteriaLimit int) bool {
		if allSegments[0].l > 0 { //first segment starts after criteria. criteria always starts at 0
			return true
		}
		//Now see if there are any other gaps in criteria range
		for i := 1; i < len(allSegments) && allSegments[i].l < criteriaLimit; i++ {
			if allSegments[i].l > allSegments[i-1].r {
				return true
			}
		}
		return false
	}

	allSegments := []Segment{{2, 10}, {4, 12}, {0, 8},
		{8, 11}, {10, 13}, {16, 20}}
	sorting.QuickSort(allSegments, lt)

	a := isThereAGap(allSegments, 17)
	if !a {
		t.Errorf("Should have found gap")
	}
	allSegments2 := []Segment{{2, 10}, {0, 3}, {5, 12}}
	sorting.QuickSort(allSegments2, lt)
	a = isThereAGap(allSegments2, 9)
	if a {
		t.Errorf("Should not have found gap")
	}
}
