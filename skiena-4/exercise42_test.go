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

type RamanujanPair struct {
	A int
	B int
	C int
	D int
}

func FindBothPairs(maybeRamanujanNumber int) option.Option[RamanujanPair] {
	f := func(p1 propcheck.Pair[int, int]) option.Option[RamanujanPair] {
		p2 := findAPair(maybeRamanujanNumber, p1)
		r := RamanujanPair{
			A: p1.A,
			B: p1.B,
		}
		return option.Map(p2, func(x propcheck.Pair[int, int]) RamanujanPair {
			return RamanujanPair{
				A: r.A,
				B: r.B,
				C: x.A,
				D: x.B,
			}
		})
	}
	ramanujanPair1 := findAPair(maybeRamanujanNumber, propcheck.Pair[int, int]{})
	return option.FlatMap[propcheck.Pair[int, int]](ramanujanPair1, f)
}

func TestIsRamanujanNumber2(t *testing.T) {
	for x := 1; x < 100000; x++ {
		steps2 = 0
		actual := FindBothPairs(x)
		switch v := actual.(type) {
		case option.None[RamanujanPair]:
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
				t.Errorf("unexpected Ramanujan number %v with cube roots%v and steps:%v", x, v, steps2)
			}

		default:
			t.Errorf("unexpected failure")
		}
	}
}
