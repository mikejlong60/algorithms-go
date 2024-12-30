package skiena_1

import (
	"fmt"
	"testing"
)

func kangaroo(x1, v1, x2, v2 int32) string {
	var result string
	if (x2 > x1) && (v2 >= v1) {
		result = "NO"
	} else if (x2-x1)%(v1-v2) == 0 {
		result = "YES"
	} else {
		result = "NO"
	}
	return result
}

func TestKangaroo(t *testing.T) {

	fmt.Println(kangaroo(0, 3, 4, 2))
	//xxs := kangaroo(3, []int32{1, 2, 3, 4, 5})
	//fmt.Print(xxs)

}
