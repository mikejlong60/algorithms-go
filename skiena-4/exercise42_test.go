package skiena_4

import (
	"fmt"
	"testing"

	"github.com/greymatter-io/golangz/option"
	"github.com/greymatter-io/golangz/propcheck"
)

var steps2 = 0

func cubed(x int) int {
	return x * x * x
}

func ramanujanNumber(aCandidate, bCandidate int, ramanujanPair propcheck.Pair[int, int], maybeRamanujanNumber int) option.Option[propcheck.Pair[int, int]] {
	steps2++
	cubedACandidate := cubed(aCandidate)
	cubedBCandidate := cubed(bCandidate)
	if cubedACandidate > maybeRamanujanNumber {
		return option.None[propcheck.Pair[int, int]]{}
	}
	if cubedACandidate+cubedBCandidate > maybeRamanujanNumber {
		return option.None[propcheck.Pair[int, int]]{}
	}

	if cubedACandidate+cubedBCandidate == maybeRamanujanNumber {
		if bCandidate == ramanujanPair.A || bCandidate == ramanujanPair.B { // You already found this one
			return ramanujanNumber(aCandidate, bCandidate+1, ramanujanPair, maybeRamanujanNumber)
		}

		return option.Some[propcheck.Pair[int, int]]{propcheck.Pair[int, int]{aCandidate, bCandidate}}

	}

	return ramanujanNumber(aCandidate, bCandidate+1, ramanujanPair, maybeRamanujanNumber)
}

func findAPair(maybeRamanujanNumber int, ramanujanPair propcheck.Pair[int, int]) option.Option[propcheck.Pair[int, int]] {
	for i := 0; i <= maybeRamanujanNumber; i++ {
		b := ramanujanNumber(i, i+1, ramanujanPair, maybeRamanujanNumber)
		switch v := b.(type) {
		case option.None[propcheck.Pair[int, int]]:
			b = ramanujanNumber(i, i+1, ramanujanPair, maybeRamanujanNumber)
		default:
			return v
		}
	}
	return option.None[propcheck.Pair[int, int]]{}
}

func TestIsRamanujanNumber(t *testing.T) {
	aa := func(x propcheck.Pair[int, int]) option.Option[propcheck.Pair[int, int]] {
		return findAPair(1729, x)
	}
	verify := func(actual option.Option[propcheck.Pair[int, int]], expected propcheck.Pair[int, int]) {
		switch v := actual.(type) {
		case option.None[propcheck.Pair[int, int]]:
			t.Errorf("unexpected none")
		case option.Some[propcheck.Pair[int, int]]:
			if !(v.Value.A == expected.A && v.Value.B == expected.B) {
				t.Errorf("expected%v but got %v", expected, v)
			}
		default:
			t.Errorf("unexpected failure")
		}

	}
	ramanujanPair1 := findAPair(1729, propcheck.Pair[int, int]{})
	ramanujanPair2 := option.FlatMap[propcheck.Pair[int, int]](ramanujanPair1, aa)
	verify(ramanujanPair1, propcheck.Pair[int, int]{1, 12})
	verify(ramanujanPair2, propcheck.Pair[int, int]{9, 10})
	fmt.Printf("steps%v\n", steps2)
}
