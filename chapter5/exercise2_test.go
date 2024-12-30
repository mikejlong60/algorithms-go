package chapter5

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestNumberOfInversionsGTDouble(t *testing.T) {
	lt := func(l, r int) bool {
		if l < r {
			return true
		} else {
			return false
		}
	}
	isInversion := func(l, r int) bool {
		if l > r*2 {
			return true
		} else {
			return false
		}
	}
	xs := []int{50, 4, 3, 2, 1}
	_, inversions := MergeSortWithInversionChecking(xs, 0, isInversion, lt)
	log.Infof("Number of inversions:%v", inversions)
	expectedInversions := 1
	if inversions != expectedInversions {
		t.Errorf("Actual:%v Expected:%v", inversions, expectedInversions)
	}
}
