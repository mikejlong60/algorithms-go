package skiena_4

import (
	"fmt"
	"testing"
)

/**  Order a random array filled with red, white and blue marbles in red,white,blue order.
 You can only look at the color and swap them.
 And do it in linear rime O(n)
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

func orderAsRedWhiteAndBlue(xs []string) {
	var i = 0
	var j = i + 1
	var currentColor = "r"
	//var nextColor = "w"

	for ; i < len(xs); i++ {
		j = i
		if xs[i] == currentColor {
			continue
		} else {
			for ; j < len(xs); j++ {
				if xs[j] == currentColor {
					xs[i], xs[j] = xs[j], xs[i]
					break
				}
			}
		}
	}
}

func TestOrderAsRedWhiteAndBlue(t *testing.T) {
	xs := []string{"r", "r", "b", "r", "w", "b", "r", "w", "r"}

	orderAsRedWhiteAndBlue(xs)
	fmt.Println(xs)
	//expected := []string{"r","r","r","r","w","w","b","b"}

	//if expected != actual {
	//	t.Errorf("Expected:[%v], Actual:[%v]", expected, actual)
	//}

}
