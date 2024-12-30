package chapter5

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestMergeSortInversion1(t *testing.T) {
	lt := func(l, r int) bool {
		if l < r {
			return true
		} else {
			return false
		}
	}
	xs := []int{2, 4, 1, 3, 5}
	isInversion := func(l, r int) bool {
		if l > r {
			return true
		} else {
			return false
		}
	}

	_, inversions := MergeSortWithInversionChecking(xs, 0, isInversion, lt)
	log.Infof("Number of inversions:%v", inversions)
	expectedInversions := 3
	if inversions != expectedInversions {
		t.Errorf("Actual:%v Expected:%v", inversions, expectedInversions)
	}
}

func TestMergeSortInversion2(t *testing.T) {
	lt := func(l, r int) bool {
		if l < r {
			return true
		} else {
			return false
		}
	}
	isInversion := func(l, r int) bool {
		if l > r {
			return true
		} else {
			return false
		}
	}
	xs := []int{5, 4, 3, 2, 1}
	_, inversions := MergeSortWithInversionChecking(xs, 0, isInversion, lt)
	log.Infof("Number of inversions:%v", inversions)
	expectedInversions := 2
	if inversions != expectedInversions {
		t.Errorf("Actual:%v Expected:%v", inversions, expectedInversions)
	}
}
