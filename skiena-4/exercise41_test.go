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
	rng := propcheck.SimpleRNG{976542023}

	res := propcheck.ChooseArray(100000, 100000, propcheck.ChooseInt(-100000, 100000))
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
	propcheck.ExpectSuccess[[]int](t, test.Run(propcheck.RunParms{100, rng}))
}

func TestBinarySearchNoFind(t *testing.T) {
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}

	res := propcheck.ChooseArray(1, 5000, propcheck.ChooseInt(0, 100000))
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
	propcheck.ExpectSuccess[[]int](t, test.Run(propcheck.RunParms{100, rng}))
}

func TestBinarySearchEmptyArray(t *testing.T) {
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}

	res := propcheck.ChooseArray(0, 0, propcheck.ChooseInt(0, 100000))
	sortIt := func(xs []int) bool {
		steps1 = 0
		fmt.Printf("Generated array of length:%v\n", len(xs))
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
	propcheck.ExpectSuccess[[]int](t, test.Run(propcheck.RunParms{100, rng}))
}

func TestTwoApproaches(t *testing.T) {
	rng := propcheck.SimpleRNG{976542023}

	res := propcheck.ChooseArray(100000, 100000, propcheck.ChooseInt(-100000, 100000))
	sortIt := func(xs []int) bool {
		steps1 = 0
		fmt.Printf("Generated array of length:%v\n", len(xs))
		lookfor := xs[len(xs)/2]
		sort.Ints(xs)

		cutoff := float32(len(xs)) * .6
		goodCustomers := xs[:int(cutoff)]
		notGoodCustomers := xs[int(cutoff):]
		start := time.Now()
		answer := binarySearch(xs, lookfor)
		fmt.Printf("binary of search of whole array took:%v\n", time.Since(start))
		start = time.Now()
		answer = binarySearch(goodCustomers, lookfor)
		if !answer {
			answer = binarySearch(notGoodCustomers, lookfor)
		}
		fmt.Printf("binary of search using special split algo took:%v\n", time.Since(start))

		return answer
	}
	verifySuccess := func(actual bool) (bool, error) {
		if !actual {
			return false, fmt.Errorf("expected %v, got %v", true, actual)
		}
		return true, nil
	}
	test := propcheck.ForAll(res, "Binary search an array of ints.", sortIt, verifySuccess)
	propcheck.ExpectSuccess[[]int](t, test.Run(propcheck.RunParms{100, rng}))
}
