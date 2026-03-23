package skiena_4

import (
	"testing"
)

// For an n x n array of integers in row strictly increasing order and column strictly decreasing order
// write an efficient algorithm that counts the number of zeros in the array.

// Answer: zeros are always on a diagonal, increasing to right as you increase y-axis from wherever
// you find the first zero:

// -1 0 1 2
//	0 1 2 4
//	1 2 3 6

func CountZeros(xs [][]int) int {
	count := 0
	for i := 0; i < len(xs); i++ {
		j := 0

		// As soon as you find a gap in the diagonal no more zeros are possible
		for j < len(xs[i]) {
			if xs[i][j] == 0 {
				count++
				j++
				break
			}
			j++
		}
	}
	return count
}

func TestZeroCount(t *testing.T) {
	xs := [][]int{{1, 2, 3, 6}, {0, 1, 2, 4}, {-1, 0, 1, 2}}
	count := CountZeros(xs)
	if count != 2 {
		t.Errorf("expected: 2 got: %v", count)
	}
}
