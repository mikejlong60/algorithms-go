package skiena_4

import (
	"github.com/greymatter-io/golangz/arrays"
	"github.com/mikejlong60/algorithms-go/chapter5"
	"testing"
)

func TestUnfairestTeams(t *testing.T) {
	lt := func(l, r int) bool {
		if l < r {
			return true
		} else {
			return false
		}
	}
	eq := func(l, r int) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}
	r := []int{1, 1, 13, 14, 234, 43, 13, 45, 15, 26, 7, 18, 90, 1, 12, 23}
	rr := chapter5.MergeSort(r, lt)
	worstTeam := rr[0 : len(r)/2]
	bestTeam := rr[len(r)/2 : len(r)]
	expectedWorst := []int{1, 1, 1, 7, 12, 13, 13, 14}
	expectedBest := []int{15, 18, 23, 26, 43, 45, 90, 234}
	if !arrays.ArrayEquality(worstTeam, expectedWorst, eq) {
		t.Errorf("\nActual worst:  %v\nExpected worst:%v", worstTeam, expectedWorst)
	}
	if !arrays.ArrayEquality(bestTeam, expectedBest, eq) {
		t.Errorf("\nActual best:  %v\nExpected best:%v", bestTeam, expectedBest)
	}
}
