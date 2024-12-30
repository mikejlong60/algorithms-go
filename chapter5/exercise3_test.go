package chapter5

import (
	"testing"
)

func TestHalfAreEquivalent(t *testing.T) {
	lt := func(l, r int) bool {
		if l < r {
			return true
		} else {
			return false
		}
	}
	eq := func(l, r int) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}
	xs := []int{3, 1, 2, 3, 4, 1, 5, 5, 12, 13, 14, 2, 2, 12, 13}
	atLeastHalfEquivalent := NumberOfEquivalences(xs, eq, lt)
	if !atLeastHalfEquivalent {
		t.Errorf("Wrong")
	}
}

func TestHalfAreNotEquivalent(t *testing.T) {
	lt := func(l, r int) bool {
		if l < r {
			return true
		} else {
			return false
		}
	}
	eq := func(l, r int) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}
	xs := []int{3, 1, 2, 3, 4, 1, 5}
	halfAreEquivalent := NumberOfEquivalences(xs, eq, lt)
	if halfAreEquivalent {
		t.Errorf("Wrong")
	}
}
