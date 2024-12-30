package skiena_1

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"testing"
	"time"
)

func left_lt_middle_middle_gt_right(l, m, r int) bool {
	if l < m && m > r {
		return true
	} else {
		return false
	}
}
func left_gt_middle_middle_lt_right(l, m, r int) bool {
	if l > m && m < r {
		return true
	} else {
		return false
	}
}

func oddNumber(i int) bool {
	if (i+1)%2 != 0 {
		return true
	} else {
		return false
	}
}

func swap(xs []int, i, j int) []int {
	s1 := xs[i]
	s2 := xs[j]
	xs[j] = s1
	xs[i] = s2
	return xs
}

// Given an integer array nums, reorder it such that nums[0] < nums[1] > nums[2] < nums[3]....
//
// You may assume the input array always has a valid answer.
//
// Input: nums = [1,5,1,1,6,4]
// Output: [1,6,1,5,1,4]
// Explanation: [1,4,1,5,1,6] is also accepted.
//
// Input: nums = [1,3,2,2,3,1]
// Output: [2,3,1,3,1,2]

// Not a pure function
func wiggleSort(xs []int) []int {
	findNextHigher := func(xs []int, numberToCompare, startingIdx int) int {
		for i := startingIdx; i < len(xs); i++ { //TODO This loop only goes to end from startingIdx. If you get to end search from beginning until you reach startingIdx. And make sure you don't disturb low-high-low or high-low-high rule
			if numberToCompare < xs[i] {
				return i
			}
		}
		return -1 //TODO temporary until you do above TODO
	}
	findNextLower := func(xs []int, numberToCompare, startingIdx int) int {
		for i := startingIdx; i < len(xs); i++ { //TODO This loop only goes to end from startingIdx. If you get to end search from beginning until you reach startingIdx. And make sure you don't disturb low-high-low or high-low-high rule
			if numberToCompare > xs[i] {
				return i
			}
		}
		return -1 //TODO temporary until you do above TODO
	}
	swap2And3OrFurtherAhead := func(xs []int, i int) []int {
		if oddNumber(i) {
			if xs[i+2] > xs[i+1] {
				xs = swap(xs, i+2, i+1)
			} else {
				xs = swap(xs, i+1, findNextLower(xs, xs[i+1], i+3))
			}
		} else {
			if xs[i+2] > xs[i+1] {
				xs = swap(xs, i+2, i+1)
			} else {
				xs = swap(xs, i+1, findNextHigher(xs, xs[i+1], i+3))
			}
		}
		return xs
	}

	for i := 0; i < len(xs)-2; i++ {
		if oddNumber(i) { //an odd element number
			correctOrder := left_lt_middle_middle_gt_right(xs[i], xs[i+1], xs[i+2])
			if !correctOrder {
				xs = swap2And3OrFurtherAhead(xs, i)
			}
		} else { // an even element number
			correctOrder := left_gt_middle_middle_lt_right(xs[i], xs[i+1], xs[i+2])
			if !correctOrder {
				xs = swap2And3OrFurtherAhead(xs, i)
			}
		}
	}
	return xs
}

func TestWiggleSort(t *testing.T) {
	g0 := propcheck.ArrayOfN(3, propcheck.ChooseInt(0, 500))
	g1 := propcheck.ArrayOfN(3, propcheck.ChooseInt(501, 1000))
	g3 := propcheck.Map2(g0, g1, func(a, b []int) []int {
		return append(a, b...)
	})

	g4 := func(xs []int) func(propcheck.SimpleRNG) ([]int, propcheck.SimpleRNG) {
		a := propcheck.ArrayOfN(len(xs), propcheck.ChooseInt(0, len(xs)-1))
		b := propcheck.ArrayOfN(len(xs), propcheck.ChooseInt(0, len(xs)-1))
		r := propcheck.Map2(a, b, func(ab, cd []int) []int {
			for i, _ := range xs { //shuffle array xs
				idx1 := ab[i]
				idx2 := cd[i]
				//swap elements in xs array
				xs = swap(xs, idx2, idx1)
			}
			return xs
		})
		return r
	}

	g5 := propcheck.FlatMap(g3, g4)

	//g55 := propcheck.Id([]int{1, 1, 1, 2, 2, 2})
	now := time.Now().Nanosecond()
	rng := propcheck.SimpleRNG{now}

	verify := func(xs []int) bool {
		var i = 0
		var r bool

		for {
			if i+2 < len(xs) {
				if oddNumber(i) { //an odd element number
					r = left_lt_middle_middle_gt_right(xs[i], xs[i+1], xs[i+2])
					if !r {
						break
					}
				} else { // an even element number
					r = left_gt_middle_middle_lt_right(xs[i], xs[i+1], xs[i+2])
					if !r {
						break
					}
				}
			} else {
				break
			}
			i = i + 1
		}
		return r
	}
	prop := propcheck.ForAll(g5,
		"Verify wiggle sort  \n",
		func(xs []int) []int {
			wiggleSort(xs)
			return xs
		},
		func(xs []int) (bool, error) {
			var errors error

			actual := wiggleSort(xs)

			if verify(actual) {
				fmt.Println("Correct!!!")
			} else {
				errors = fmt.Errorf("Actual:%v", xs)
			}
			if errors != nil {
				return false, errors
			} else {
				return true, nil
			}
		},
	)
	result := prop.Run(propcheck.RunParms{100, rng})
	propcheck.ExpectSuccess[[]int](t, result)

}
