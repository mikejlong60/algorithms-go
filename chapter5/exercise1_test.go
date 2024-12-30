package chapter5

import (
	"testing"
)

func TestFindMedianOfTwoSets1(t *testing.T) {
	a := []float32{1.0, 2.0, 3.0}
	b := []float32{4.0, 5.0, 6.0}

	actual := MedianOfTwoLists(a, b)
	expected := (float32)(3.5)

	if actual != expected {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
}

func TestFindMedianOfTwoSets2(t *testing.T) {
	a := []float32{12.0, 13.0, 1.0, 2.0}
	b := []float32{300.0, 0.0, 3.0, 100.0}

	actual := MedianOfTwoLists(a, b)
	expected := (float32)(7.5)

	if actual != expected {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
}
