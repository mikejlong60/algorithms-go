package skiena_4

import (
	"fmt"
	"testing"

	"github.com/greymatter-io/golangz/option"
)

// For an n x n array of integers in row strictly increasing order and column strictly decreasing order
// write an efficient algorithm that counts the number of zeros in the array.

// Answer:
//  x = 0 -- current row
//  A = 0 -- zero count
//  B = 0 -- excluded columns
//    Starting at  row x go across it, excluding all columns in array B.
//    If you find a zero increment A, the zero count and append the column number to an array b.
//    Repeat until you have addressed all rows
//   Return A

// Real answer - zeros are always on a diagonal increasing to right as you increase y axis from wherever
// you find the first zero
// -1 0 1
//
//	0 1 2
//	1 2 3
func TestZeroCount(t *testing.T) {
	for x := 1; x < 200000; x++ {
		steps2 = 0
		actual := FindBothPairs(x)
		switch v := actual.(type) {
		case option.None[RamanujanPair]:
			fmt.Printf("not a Ramanujan number:%v  in %v steps\n", x, steps2)
		case option.Some[RamanujanPair]:
			if x == 1729 {
				fmt.Printf("actual:%v steps:%v\n", actual, steps2)
				if !(v.Value.A == 1 && v.Value.B == 12 && v.Value.C == 9 && v.Value.D == 10) {
					t.Errorf("expected{{1 12 9 10}} for %v but got %v and took %v steps", x, v, steps2)
				}
			} else if x == 4104 {
				fmt.Printf("actual:%v steps:%v\n", actual, steps2)
				if !(v.Value.A == 2 && v.Value.B == 16 && v.Value.C == 9 && v.Value.D == 15) {
					t.Errorf("expected{{2, 16, 9, 15}} for %v but got %v and took %v steps", x, v, steps2)
				}
			} else {
				fmt.Printf("another Ramanujan number:%v with cube roots:%v and steps:%v\n", x, v, steps2)
			}

		default:
			t.Errorf("unexpected failure")
		}
	}
}
