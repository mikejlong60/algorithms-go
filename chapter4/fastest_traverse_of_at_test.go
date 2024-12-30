package chapter4

import (
	"github.com/greymatter-io/golangz/arrays"
	"testing"
)

// This is solved exercise 1
func FastestATTraverse(maxDistancePerDay, lastStopIdx int, campSites, route []int) (int, int, []int, []int) {
	if len(route) == 0 {
		route = append(route, campSites[0])
		lastStopIdx = lastStopIdx + 1
		return FastestATTraverse(maxDistancePerDay, lastStopIdx, campSites, route)
	} else if lastStopIdx > len(campSites)-1 {
		route = append(route, campSites[len(campSites)-1])
		return maxDistancePerDay, lastStopIdx, campSites, route
	} else if campSites[lastStopIdx]-route[len(route)-1] > maxDistancePerDay {
		lastStopIdx = lastStopIdx - 1
		route = append(route, campSites[lastStopIdx])
		return FastestATTraverse(maxDistancePerDay, lastStopIdx, campSites, route)
	} else {
		lastStopIdx = lastStopIdx + 1
		return FastestATTraverse(maxDistancePerDay, lastStopIdx, campSites, route)
	}
}
func TestFastestTraverseOfAT1(t *testing.T) {
	eq := func(l, r int) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}
	campSites := []int{1, 4, 6, 8, 10, 12, 14, 18, 21, 22, 25, 30}
	maxDistancePerDay := 5
	_, _, _, actual := FastestATTraverse(maxDistancePerDay, 0, campSites, []int{})
	expected := []int{1, 6, 10, 14, 18, 22, 25, 30}
	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Expected:%v Actual:%v", expected, actual)
	}
}

func TestFastestTraverseOfAT2(t *testing.T) {
	eq := func(l, r int) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}
	campSites := []int{1, 4, 6, 8, 10, 12, 14, 18, 21, 22, 25, 30}

	maxDistancePerDay := 7
	_, _, _, actual := FastestATTraverse(maxDistancePerDay, 0, campSites, []int{})
	expected := []int{1, 8, 14, 21, 25, 30}
	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Expected:%v Actual:%v", expected, actual)
	}

}
