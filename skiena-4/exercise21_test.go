package skiena_4

import (
	"math"
	"testing"
)

//  Find the median of a set of numbers using some sort of partitioning technique like Quicksort.

func TestBigOhNMedian(t *testing.T) {
	bigOhMedian := func(mid int, xs []int) int {
		start := xs[mid]
		swap := func(x int, y int, xs []int) {
			tx := xs[x]
			xs[x] = y
			xs[y] = tx
		}

		for i := 0; i < mid; i++ {
			if start < xs[i] {
				swap(start, xs[i], xs)
			}
		}
		return xs[0]
	}

	xs := []int{0, 2, 7, 6, 11, 10, 5, 8, 9}

	m := float64(len(xs) / 2)
	mid :=int(math.Ceil(m))

	var actual = bigOhMedian(mid, xs)

	expected := 8
	if expected != actual {
		t.Errorf("Expected:[%v], Actual:[%v]", expected, actual)
	}

}
