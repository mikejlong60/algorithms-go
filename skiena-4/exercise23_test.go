package skiena_4

import (
	"fmt"
	"github.com/greymatter-io/golangz/arrays"
	"testing"
)

/**  Order a random array filled with red, white and blue marbles in red,white,blue order.
 You can only look at the color and swap them.
 And do it in linear time O(n)
Algorithm:
Start at the first element.
  If it is Red, move to next element
  Is the next element red?
    Yes - move i to next element
    No - move i to that element.
    Iterate until you hit a red or reach the end of the array.
    If you hit a red, swap that element with i.
  end

  At this point start at the end of the reds and start looking for white elements.
  Is next element white?
     yes - move to next element
     no - move to next element
     Iterate until you hit a white or reach the end of the array
  end
  done
*/

var steps = 0

func orderAsRedWhiteAndBlue(xs []string, i int, currentColor string) int {
	var j = 0
	for ; i < len(xs); i++ {
		if j == len(xs) {
			return i - 1
		}
		steps++
		if xs[i] == currentColor {
			continue
		} else {
			j = i
			for ; j < len(xs); j++ {
				steps++
				if xs[j] == currentColor {
					//matchedCurrentColor = true
					xs[i], xs[j] = xs[j], xs[i]
					if j < len(xs) {
						break
					}
				}
			}
		}
	}
	return i
}

func TestOrderAsRedWhiteAndBlue(t *testing.T) {

	eq := func(l, r string) bool {
		if l == r {
			return true
		}
		return false
	}

	xs := []string{"r", "r", "b", "r", "w", "b", "r", "r", "b", "r", "w", "b"}

	i := orderAsRedWhiteAndBlue(xs, 0, "r")
	orderAsRedWhiteAndBlue(xs, i, "w")
	fmt.Println(steps)
	fmt.Println(xs)
	expected := []string{"r", "r", "r", "r", "r", "r", "w", "w", "b", "b", "b", "b"}

	if !arrays.ArrayEquality(xs, expected, eq) {
		t.Errorf("Expected:[%v], Actual:[%v]", expected, xs)
	}

	ys := []string{"r", "w", "b", "b", "w", "b", "r"}
	steps = 0
	k := orderAsRedWhiteAndBlue(ys, 0, "r")
	orderAsRedWhiteAndBlue(ys, k, "w")
	fmt.Println(steps)
	expectedys := []string{"r", "r", "w", "w", "b", "b", "b"}
	if !arrays.ArrayEquality(ys, expectedys, eq) {
		t.Errorf("Expected:[%v], Actual:[%v]", expectedys, ys)
	}
}
