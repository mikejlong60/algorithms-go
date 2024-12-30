package skiena_4

import (
	"github.com/greymatter-io/golangz/arrays"
	"github.com/greymatter-io/golangz/propcheck"
	"math"
	"testing"
)

// Compute the mode of a list of numbers.
func TestMode(t *testing.T) {
	//Big O(n) with a hashmap

	in := []int{4, 6, 2, 4, 3, 1}

	g := func(modeAccum propcheck.Pair[int, map[int]int], currentNKey int) propcheck.Pair[int, map[int]int] {
		_, currentNExists := modeAccum.B[currentNKey]
		currentMode, currentModeExists := modeAccum.B[modeAccum.A]
		if !currentNExists && !currentModeExists {
			modeAccum.B[currentNKey] = 1
			modeAccum.A = currentNKey
		} else if !currentNExists && currentModeExists {
			modeAccum.B[currentNKey] = 1
		} else if currentNExists && currentModeExists {
			nMode := modeAccum.B[currentNKey] + 1
			modeAccum.B[currentNKey] = nMode
			if nMode > currentMode {
				modeAccum.A = currentNKey
			}
		}
		return modeAccum
	}

	mode := arrays.FoldLeft(in, propcheck.Pair[int, map[int]int]{math.MinInt, map[int]int{}}, g)
	if mode.A != 4 {
		t.Error("Expected", 4, "got", mode.A)
	}
	if mode.B[4] != 2 {
		t.Error("Expected", 2, "got", mode.B[4])
	}
}
