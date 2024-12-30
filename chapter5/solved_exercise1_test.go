package chapter5

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/hashicorp/go-multierror"
	"sort"
	"testing"
	"time"
)

var max = func(a, b int) bool {
	if a > b {
		return true
	} else {
		return false
	}
}

func maxIntOrEmpty(xs []int) []int {
	if len(xs) == 3 {
		if max(xs[2], xs[0]) && max(xs[2], xs[1]) { //peak is last
			return []int{xs[2]}
		} else if max(xs[1], xs[0]) && max(xs[1], xs[2]) { //peak is second
			return []int{xs[1]}
		} else if max(xs[0], xs[1]) && max(xs[0], xs[2]) { //peak is first
			return []int{xs[0]}
		} else { //they are all equal so just return first
			return []int{xs[0]}
		}
	} else if len(xs) == 2 {
		if max(xs[0], xs[1]) {
			return []int{xs[0]}
		} else {
			return []int{xs[1]}
		}
	} else if len(xs) == 1 {
		return xs
	} else if len(xs) == 0 {
		return xs
	} else {
		return []int{}
	}
}

func TestMaxInt(t *testing.T) {
	g0 := propcheck.ChooseArray(1000, 50000, propcheck.ChooseInt(0, 10000000))

	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	prop := propcheck.ForAll(g0,
		"Exercise 1a, the peak value of an array of ints.",
		func(xs []int) []int {
			return xs
		},
		func(xs []int) (bool, error) {
			dest := make([]int, len(xs))
			copy(dest, xs)
			st1 := time.Now()
			b := PeakOfAList(xs, maxIntOrEmpty)
			var errors error
			st2 := time.Now()
			sort.Ints(dest)
			bb := len(dest)
			if bb > 0 {
				bbb := dest[bb-1]
				fmt.Printf("Go max took:%v max is:%v\n", time.Since(st2), bbb)
			}
			fmt.Printf("my max took:%v max is:%v\n", time.Since(st1), b)

			l := len(dest)
			if l > 0 {
				if dest[l-1] != b[0] {
					errors = multierror.Append(errors, fmt.Errorf("Expected max to be be:%v but was:%v", dest[l-1], b[0]))
				}
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
