package skiena_4

import (
	"fmt"
	"testing"
)

func isRamanujanNumber(firstCubedNumber, bCandidate int, maybeRamanujanNumber int) bool {
	cubed := func(x int) int {
		return x * x * x
	}

	if firstCubedNumber > maybeRamanujanNumber {
		return false
	}
	if firstCubedNumber+cubed(bCandidate) > maybeRamanujanNumber {
		return false
	}

	if firstCubedNumber+cubed(bCandidate) == maybeRamanujanNumber {
		return true
	}

	return isRamanujanNumber(firstCubedNumber, bCandidate+1, maybeRamanujanNumber)
}

func TestIsRamanujanNumber(t *testing.T) {

	isRamanujanNumber := isRamanujanNumber(0, 1, 1729)

	fmt.Println(isRamanujanNumber)
}
