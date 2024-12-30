package skiena_4

import (
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/propcheck"
	"testing"
)

func TestNestedSort(t *testing.T) {
	//Big O(n) because already sorted enough

	eq := func(l, r propcheck.Pair[int, string]) bool {
		if l.A == r.A && l.B == r.B {
			return true
		} else {
			return false
		}
	}
	in := []propcheck.Pair[int, string]{{1, "blue"}, {3, "red"}, {4, "blue"}, {6, "yellow"}, {9, "red"}}

	redBucket := make([]propcheck.Pair[int, string], 0)
	blueBucket := make([]propcheck.Pair[int, string], 0)
	yellowBucket := make([]propcheck.Pair[int, string], 0)

	for _, cp := range in {
		switch cp.B {
		case "red":
			redBucket = append(redBucket, cp)
		case "blue":
			blueBucket = append(blueBucket, cp)
		case "yellow":
			yellowBucket = append(yellowBucket, cp)
		}
	}
	actual := append(append(redBucket, blueBucket...), yellowBucket...)
	expected := []propcheck.Pair[int, string]{{3, "red"}, {9, "red"}, {1, "blue"}, {4, "blue"}, {6, "yellow"}}
	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
}
