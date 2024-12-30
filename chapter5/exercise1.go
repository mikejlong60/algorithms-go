package chapter5

import (
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/sorting"
)

func MedianOfTwoLists(l, r []float32) float32 {

	//Append the two lists
	//Then sort them
	//Then take their median
	lt := func(l, r float32) bool {
		if l < r {
			return true
		} else {
			return false
		}
	}

	medianOfSet := func(xs []float32) float32 {
		sorting.QuickSort(xs, lt)
		size := len(xs)
		if size%2 == 0 {
			ll := xs[(size/2)-1]
			rr := xs[(size / 2)]
			return (ll + rr) / 2
		} else {
			return xs[size/2+1]
		}
	}
	return medianOfSet(arrays.Append(l, r))

}
