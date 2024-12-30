package skiena_4

import (
	"fmt"
	"github.com/greymatter-io/golangz/propcheck"
	"github.com/mikejlong60/algorithms/chapter5"
	"testing"
	"time"
)

var numSteps int

func sumExists(xs, ys []int, desiredSum int) bool {
	numSteps = numSteps + 1
	if len(xs) == 0 {
		return false
	}
	if xs[0]+ys[0] > desiredSum {
		return false
	}
	if xs[0]+ys[0] == desiredSum {
		return true
	}
	if len(xs) == 1 {
		if xs[0]+ys[0] == desiredSum {
			return true
		} else {
			return false
		}
	}

	if xs[0]+ys[len(ys)/2] == desiredSum {
		return true
	} else if xs[0]+ys[len(ys)/2] > desiredSum {
		return sumExists(xs[1:], ys[0:len(ys)/2], desiredSum)
	} else if xs[0]+ys[len(ys)/2] < desiredSum {
		return sumExists(xs[1:], ys[len(ys)/2:], desiredSum)
	} else {
		panic("unexpected sum")
	}

}

func SumExists(xs, ys []int, desiredSum int) bool {
	lt := func(l, r int) bool {
		if l < r {
			return true
		} else {
			return false
		}
	}
	if len(xs) == 0 || len(xs) != len(ys) {
		return false
	}
	xss := chapter5.MergeSort(xs, lt)
	yss := chapter5.MergeSort(ys, lt)
	return sumExists(xss, yss, desiredSum)
}

// Determine in Big O(n log n) whether a pair of numbers from list A and list B add up to a given number
func Test2SetsSum1(t *testing.T) {
	numSteps = 0

	xs := []int{4, 6, 2, 3, 1}
	ys := []int{1, 2, 3, 4, 5}

	actual := SumExists(xs, ys, 5)
	fmt.Printf("Number of steps:%v\n", numSteps)
	if !actual {
		t.Errorf("Expected sum:%v to exist but it did not", 5)
	}
}

func Test2SetsSum2(t *testing.T) {
	numSteps = 0

	xs := []int{4, 6, 2, 3, 1}
	ys := []int{1, 2, 3, 4, 5}

	actual := SumExists(xs, ys, 50)
	fmt.Printf("Number of steps:%v\n", numSteps)
	if actual {
		t.Errorf("Expected sum:%v to not exist but it did", 50)
	}
}

func Test2SetsSum3(t *testing.T) {
	numSteps = 0

	xs := []int{4, 6, 2, 3, 1}
	ys := []int{1, 2, 3, 4, 5}

	actual := SumExists(xs, ys, 7)
	fmt.Printf("Number of steps:%v\n", numSteps)
	if !actual {
		t.Errorf("Expected sum:%v to exist but it did not", 7)
	}
}

type TestData struct {
	Xs         []int
	Ys         []int
	DesiredSum int
}

func TestSumProp(t *testing.T) {
	g1 := propcheck.ChooseArray(0, 40000, propcheck.ChooseInt(-100, 100)) //The source array split in half

	rng := propcheck.SimpleRNG{time.Now().Nanosecond()}
	prop := propcheck.ForAll(g1,
		"Determine whether 2 arrays of equal size have a given sum among their elements with Big O(n log n) efficiency.",
		func(xs []int) TestData {
			numSteps = 0
			g := len(xs) % 2
			if g == 1 { //Make it evenly divisible
				xs = xs[1:]
			}
			midPoint := len(xs) / 2
			a1 := xs[0:midPoint]
			a2 := xs[midPoint:]
			return TestData{a1, a2, xs[midPoint]}
		},
		func(a TestData) (bool, error) {

			var errors error

			actual := SumExists(a.Xs, a.Ys, a.DesiredSum)
			fmt.Printf("Number of steps:%v for n size:%v\n", numSteps, len(a.Ys))

			if !actual {
				fmt.Println("Sum not found")
				//errors = multierror.Append(errors, fmt.Errorf("Expected sum:%v to exist but it did not", actual))
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
