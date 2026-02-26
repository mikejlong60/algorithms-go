package skiena_4

import (
	"fmt"
	"sort"
	"testing"
	"time"

	"github.com/greymatter-io/golangz/propcheck"
)

var steps1 = 0

func binarySearch(sortedArray []int, target int) bool {
	steps1++
	if len(sortedArray) == 0 {
		return false
	}

	if len(sortedArray) == 1 {
		if sortedArray[0] == target {
			return true
		} else {
			return false
		}
	}

	if len(sortedArray) == 2 {
		if sortedArray[0] == target || sortedArray[1] == target {
			return true
		} else {
			return false
		}
	}

	midpoint := len(sortedArray) / 2
	if sortedArray[midpoint] == target {
		return true
	}

	if sortedArray[midpoint] < target {
		return binarySearch(sortedArray[midpoint:], target)
	}

	if sortedArray[midpoint] > target {
		return binarySearch(sortedArray[:midpoint], target)
	}
	return false
}

func TestBinarySearchSuccess(t *testing.T) {
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}

	res := propcheck.ChooseArray(5000000, 5000000, propcheck.ChooseInt(-100000, 100000))
	sortIt := func(xs []int) bool {
		steps1 = 0
		fmt.Printf("Generated array of length:%v\n", len(xs))
		sort.Ints(xs)
		answer := binarySearch(xs, xs[len(xs)/5])
		fmt.Printf("steps1:%v\n", steps1)

		return answer
	}
	verifySuccess := func(actual bool) (bool, error) {
		if !actual {
			return false, fmt.Errorf("expected %v, got %v", true, actual)
		}
		return true, nil
	}
	test := propcheck.ForAll(res, "Binary search an array of ints.", sortIt, verifySuccess)
	propcheck.ExpectSuccess[[]int](t, test.Run(propcheck.RunParms{10, rng}))
}

func TestBinarySearchNoFind(t *testing.T) {
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}

	res := propcheck.ChooseArray(0, 5000, propcheck.ChooseInt(0, 100000))
	sortIt := func(xs []int) bool {
		steps1 = 0
		fmt.Printf("Generated array of length:%v\n", len(xs))
		sort.Ints(xs)
		answer := binarySearch(xs, -100)
		fmt.Printf("steps1:%v\n", steps1)

		return answer
	}
	verifyFailure := func(actual bool) (bool, error) {
		if actual {
			return false, fmt.Errorf("expected %v, got %v", false, actual)
		}
		return true, nil
	}
	test := propcheck.ForAll(res, "Binary search an array of ints and fail to find it.", sortIt, verifyFailure)
	propcheck.ExpectSuccess[[]int](t, test.Run(propcheck.RunParms{10, rng}))
}
