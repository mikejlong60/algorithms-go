package skiena_1

import (
	"fmt"
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/propcheck"
	"testing"
	"time"
)

func dailyTemperatures(temperatures []int) []int {
	var result = []int{}
	for a, _ := range temperatures {
		var numDays = 0
		var noneGreater = true
		for i := a + 1; i < len(temperatures); i++ {
			numDays = numDays + 1
			if temperatures[i] > temperatures[a] {
				noneGreater = false
				break
			}
		}
		if noneGreater {
			result = append(result, 0)
		} else {
			result = append(result, numDays)
		}
	}
	return result
}

func TestTemperature(t *testing.T) {
	eq := func(l, r int) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}
	in := propcheck.Id([]int{73, 74, 75, 71, 69, 72, 76, 73})
	now := time.Now().Nanosecond()
	rng := propcheck.SimpleRNG{now}
	prop := propcheck.ForAll(in,
		"Validate integer division without * or / operators  \n",
		func(xs []int) []int {
			return xs
		},
		func(xs []int) (bool, error) {
			var errors error
			expected := []int{1, 1, 4, 2, 1, 1, 0, 0}
			//expected := []int{0, 0}
			actual := dailyTemperatures(xs)
			if !arrays.ArrayEquality(actual, expected, eq) {
				errors = fmt.Errorf("Actual:%v Expected:%v", actual, expected)
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

//
//
//[73,74,75,71,69,72,76,73]
//
//actuasl [3,2,1,3,3,2,0,0]
//edxpected [1,1,4,2,1,1,0,0]
