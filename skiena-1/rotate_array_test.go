package skiena_1

import (
	"github.com/greymatter-io/golangz/arrays"
	"testing"
)

func rotateLeft(d int32, arr []int32) []int32 {

	for i := int32(0); i < d; i++ {
		x := arr[0]
		arr = arr[1:]
		arr = append(arr, x)
	}
	return arr
}

func TestRotateArray(t *testing.T) {
	eq := func(l, r int32) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}

	actual := rotateLeft(3, []int32{1, 2, 3, 4, 5})
	expected := []int32{4, 5, 1, 2, 3}
	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Actual%v, Expected:%v\n", actual, expected)
	}
}
