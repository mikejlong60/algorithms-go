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

// Answer.
// Average time binary Search:946ns
// Average time special binary Search:1.929µs
// Average time sequential Search:149.239µs
// The special binary search is not better. The two binary search are the same.
// Any difference is due to cache warm up or something else.
// But the sequential search is always slower.
func TestTwoApproaches(t *testing.T) {
	const testCases = 30
	var totTimeBinarySearch time.Duration
	var totTimeSpecialBinarySearch time.Duration
	var totTimeSeqSearch time.Duration
	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	f := propcheck.ChooseArray(10000, 1000000, propcheck.ChooseInt(-100000, 100000))

	sortIt := func(xs []int) bool {
		lookfor := xs[len(xs)/2]

		sort.Ints(xs)

		cutoff := int(float32(len(xs)) * .6)
		goodCustomers := xs[:cutoff]
		notGoodCustomers := xs[cutoff:]
		start := time.Now()
		answer := binarySearch(goodCustomers, lookfor)
		if !answer {
			answer = binarySearch(notGoodCustomers, lookfor)
		}
		b := time.Since(start)
		totTimeSpecialBinarySearch = totTimeSpecialBinarySearch + b
		start = time.Now()
		answer = binarySearch(xs, lookfor)
		a := time.Since(start)
		totTimeBinarySearch = totTimeBinarySearch + a
		sequentialSearch := func(xs []int) bool {
			for _, x := range xs {
				if x == lookfor {
					return true
				}
			}
			return false
		}
		start = time.Now()
		answer = sequentialSearch(xs)
		c := time.Since(start)
		totTimeSeqSearch = totTimeSeqSearch + c

		return answer
	}
	verifySuccess := func(actual bool) (bool, error) {
		if !actual {
			return false, fmt.Errorf("expected %v, got %v", true, actual)
		}
		return true, nil
	}
	test := propcheck.ForAll(f, "Binary search an array of ints.", sortIt, verifySuccess)
	propcheck.ExpectSuccess[[]int](t, test.Run(propcheck.RunParms{testCases, rng}))
	fmt.Printf("Average time binary Search:%v\n", totTimeBinarySearch/testCases)
	fmt.Printf("Average time special binary Search:%v\n", totTimeSpecialBinarySearch/testCases)
	fmt.Printf("Average time sequential Search:%v\n", totTimeSeqSearch/testCases)
}
