package skiena_3

import (
	"fmt"
	"github.com/greymatter-io/golangz/option"
	"strings"
	"testing"
)

func TestAnagramDetection(t *testing.T) {

	lt := func(l, r string) bool {
		if l < r {
			return true
		} else {
			return false
		}
	}
	eq := func(l, r string) bool {
		if l == r {
			return true
		} else {
			return false
		}
	}

	aa := "incest"
	a := strings.Split(aa, "")
	b := strings.Split("insect", "")
	btree := BinaryTree(a, lt)

	for _, c := range b {
		err := option.GetOrElse(Find(btree, c, lt, eq), fmt.Sprintf("Letter:%v not in word:%v", c, aa))
		if err == fmt.Sprintf("Letter:%v not in word:%v", c, aa) {
			t.Errorf(err)
		}
	}
	if len(a) != len(b) {
		t.Errorf("Words were not same length")
	}
}
