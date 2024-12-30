package skiena_2

import (
	"github.com/greymatter-io/golangz/arrays"
	"strconv"
	"strings"
	"testing"
)

/**
Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number
of 1's in the binary representation of i.


Example 1:

Input: n = 2
Output: [0,1,1]
Explanation:
0 --> 0
1 --> 1
2 --> 10
Example 2:

Input: n = 5
Output: [0,1,1,2,1,2]
Explanation:
0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101



*/

func countBits(n int) []int {
	r := []int{}
	for i := 0; i <= n; i++ {
		d := strconv.FormatInt(int64(i), 2)
		c := strings.Count(d, "1")
		r = append(r, c)
	}

	return r
}

func TestCountBits(t *testing.T) {
	eq := func(l, r int) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}
	xs := countBits(5)
	expected := []int{0, 1, 1, 2, 1, 2}
	if !arrays.ArrayEquality(xs, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", xs, expected)
	}

	xs = countBits(2)
	expected = []int{0, 1, 1}
	if !arrays.ArrayEquality(xs, expected, eq) {
		t.Errorf("Actual:%v, Expected:%v", xs, expected)
	}
}
