package chapter2

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/hashicorp/go-multierror"
	"testing"
	"time"
)

func TestMatrixSum(t *testing.T) {

	g0 := propcheck.ChooseInt(1, 3000)
	g1 := propcheck.ChooseArray(0, 10000, g0)
	now := time.Now().Nanosecond()
	rng := propcheck.SimpleRNG{now}
	prop := propcheck.ForAll(g1,
		"Validate exercise 2.6, a sort-of sum of a matrix  \n",
		func(xs []int) [][]int64 {
			start := time.Now()
			r := sum(xs)
			fmt.Printf("Summing algorithm for an array of length:%v took:%v\n", len(xs), time.Since(start))
			return r
		},
		func(xss [][]int64) (bool, error) {
			//fmt.Println(xss)
			var errors error
			for i := 1; i < len(xss); i++ {
				last := xss[i-1][1]

				if xss[i][1] < last {
					errors = multierror.Append(errors, fmt.Errorf("Array element sum[%v] should not have been less than previous accumulated value", xss[i][1]))
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

func TestMatrixSumWithoutInnerLoop(t *testing.T) {

	g0 := propcheck.ChooseInt(1, 3000)
	g1 := propcheck.ChooseArray(0, 10000, g0)
	now := time.Now().Nanosecond()
	rng := propcheck.SimpleRNG{now}
	prop := propcheck.ForAll(g1,
		"Validate exercise 2.6, a sort-of sum of a matrix. This is more efficient that the previous version  \n",
		func(xs []int) [][]int64 {
			start := time.Now()
			r := matrixSumWithoutInnerLoop(xs)
			fmt.Printf("Summing algorithm with only 1 loop for an array of length:%v took:%v\n", len(xs), time.Since(start))
			return r
		},
		func(xss [][]int64) (bool, error) {
			//fmt.Println(xss)
			var errors error
			for i := 1; i < len(xss); i++ {
				last := xss[i-1][1]

				if xss[i][1] < last {
					errors = multierror.Append(errors, fmt.Errorf("Array element sum[%v] should not have been less than previous accumulated value", xss[i][1]))
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

func TestFoldRightMatrixSum(t *testing.T) { //This reverses the order because foldRight does that. There is a Reverse function in Golangz if that matters to you.

	g0 := propcheck.ChooseInt(1, 3000)
	g1 := propcheck.ChooseArray(0, 10000, g0)
	now := time.Now().Nanosecond()
	rng := propcheck.SimpleRNG{now}
	prop := propcheck.ForAll(g1,
		"Validate exercise 2.6, a sort-of sum of a matrix. This is more efficient that the previous version  \n",
		func(xs []int) [][]int64 {
			start := time.Now()
			r := matrixSumWithFoldRight(xs)
			fmt.Printf("Summing algorithm with FoldRight for an array of length:%v took:%v\n", len(xs), time.Since(start))
			return r
		},
		func(xss [][]int64) (bool, error) {
			//fmt.Println(xss)
			var errors error
			for i := 1; i < len(xss); i++ {
				last := xss[i-1][1]

				if xss[i][1] < last {
					errors = multierror.Append(errors, fmt.Errorf("Array element sum[%v] should not have been less than previous accumulated value", xss[i][1]))
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

func TestHolidaySongs(t *testing.T) { //This reverses the order because foldRight does that. There is a Reverse function in Golangz if that matters to you.
	g0 := propcheck.String(10)
	g1 := propcheck.ChooseArray(0, 1000, g0)
	now := time.Now().Nanosecond()
	rng := propcheck.SimpleRNG{now}
	prop := propcheck.ForAll(g1,
		"Validate exercise 2.6, a sort-of sum of a matrix. This is more efficient that the previous version  \n",
		func(xs []string) [][]string {
			start := time.Now()
			concat := func(x, y string) string {
				return fmt.Sprintf("%v --- %v", y, x)
			}
			r := matrixSumWithFoldLeft(xs, concat)
			fmt.Printf("Summing algorithm with FoldLeft for an array of length:%v took:%v\n", len(xs), time.Since(start))
			return r
		},
		func(xss [][]string) (bool, error) {
			var errors error
			for i := 1; i < len(xss); i++ {
				last := len(xss[i-1][1])

				if len(xss[i][1]) < last {
					errors = multierror.Append(errors, fmt.Errorf("Array element sum[%v] should not have been less than previous accumulated value", xss[i][1]))
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
	propcheck.ExpectSuccess[[]string](t, result)
}
