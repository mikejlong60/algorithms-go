package skiena_1

import (
	"slices"
	"testing"
)

/*
* Complete the 'hackerlandRadioTransmitters' function below.
*
* The function is expected to return an INTEGER.
* The function accepts following parameters:
*  1. INTEGER_ARRAY x
*  2. INTEGER k
 */
func hackerlandRadioTransmitters(xs []int32, k int32) int32 {
	slices.Sort(xs)
	numAntennas := 1
	startingPoint := 0
	if len(xs) == 0 {
		return 0
	}
	for a, _ := range xs {
		if a-startingPoint == 1 && xs[a]-xs[startingPoint] > k { //Neighbor from start cannot be out of range
			startingPoint = a
			numAntennas = numAntennas + 1
		} else if a-startingPoint > 1 && xs[a]-xs[startingPoint] > (k*2) { //After neighbor then entire neighborhood must be in range from start
			startingPoint = a
			numAntennas = numAntennas + 1
		}
	}
	return int32(numAntennas)
}

func TestCalculateHowManyRadioTransmittersNeeded1(t *testing.T) {
	actual := hackerlandRadioTransmitters([]int32{1, 4, 7}, 2)
	expected := int32(3)
	if actual != expected {
		t.Errorf("Actual%v, Expected:%v\n", actual, expected)
	}
}

func TestCalculateHowManyRadioTransmittersNeeded2(t *testing.T) {
	actual := hackerlandRadioTransmitters([]int32{1}, 2)
	expected := int32(1)
	if actual != expected {
		t.Errorf("Actual%v, Expected:%v\n", actual, expected)
	}
}

func TestCalculateHowManyRadioTransmittersNeeded3(t *testing.T) {
	actual := hackerlandRadioTransmitters([]int32{1, 2, 4}, 2)
	expected := int32(1)
	if actual != expected {
		t.Errorf("Actual%v, Expected:%v\n", actual, expected)
	}
}

func TestCalculateHowManyRadioTransmittersNeeded4(t *testing.T) {
	actual := hackerlandRadioTransmitters([]int32{1, 2, 4, 6, 8, 10}, 2)
	expected := int32(2)
	if actual != expected {
		t.Errorf("Actual%v, Expected:%v\n", actual, expected)
	}
}

func TestCalculateHowManyRadioTransmittersNeeded5(t *testing.T) {
	actual := hackerlandRadioTransmitters([]int32{1, 2, 4, 6, 8, 10, 20}, 2)
	expected := int32(3)
	if actual != expected {
		t.Errorf("Actual%v, Expected:%v\n", actual, expected)
	}
}

func TestCalculateHowManyRadioTransmittersNeeded6(t *testing.T) {
	actual := hackerlandRadioTransmitters([]int32{}, 2)
	expected := int32(0)
	if actual != expected {
		t.Errorf("Actual%v, Expected:%v\n", actual, expected)
	}
}

// /////////
func TestCalculateHowManyRadioTransmittersNeeded7(t *testing.T) {
	actual := hackerlandRadioTransmitters([]int32{1, 4, 7}, 3)
	expected := int32(1)
	if actual != expected {
		t.Errorf("Actual%v, Expected:%v\n", actual, expected)
	}
}

func TestCalculateHowManyRadioTransmittersNeeded8(t *testing.T) {
	actual := hackerlandRadioTransmitters([]int32{1}, 3)
	expected := int32(1)
	if actual != expected {
		t.Errorf("Actual%v, Expected:%v\n", actual, expected)
	}
}

func TestCalculateHowManyRadioTransmittersNeeded9(t *testing.T) {
	actual := hackerlandRadioTransmitters([]int32{1, 2, 4}, 3)
	expected := int32(1)
	if actual != expected {
		t.Errorf("Actual%v, Expected:%v\n", actual, expected)
	}
}

func TestCalculateHowManyRadioTransmittersNeeded10(t *testing.T) {
	actual := hackerlandRadioTransmitters([]int32{1, 2, 4, 6, 8, 10}, 3)
	expected := int32(2)
	if actual != expected {
		t.Errorf("Actual%v, Expected:%v\n", actual, expected)
	}
}

func TestCalculateHowManyRadioTransmittersNeeded11(t *testing.T) {
	actual := hackerlandRadioTransmitters([]int32{1, 2, 4, 6, 8, 10, 20}, 3)
	expected := int32(3)
	if actual != expected {
		t.Errorf("Actual%v, Expected:%v\n", actual, expected)
	}
}

func TestCalculateHowManyRadioTransmittersNeeded12(t *testing.T) {
	actual := hackerlandRadioTransmitters([]int32{1, 2, 4, 6, 8, 10, 20}, 30)
	expected := int32(1)
	if actual != expected {
		t.Errorf("Actual%v, Expected:%v\n", actual, expected)
	}
}

func TestCalculateHowManyRadioTransmittersNeeded13(t *testing.T) {
	actual := hackerlandRadioTransmitters([]int32{7, 2, 4, 6, 5, 9, 12, 11}, 2)
	expected := int32(3)
	if actual != expected {
		t.Errorf("Actual%v, Expected:%v\n", actual, expected)
	}
}
