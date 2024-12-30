package skiena_2

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func removeKDigits(num string, k int) string {
	min := func(xs []int) int {
		m := math.MaxInt
		for _, y := range xs {
			if y < m {
				m = y
			}
		}
		return m
	}

	split := func(num string, k int) []int {
		r := []int{}
		a := 0
		b := a + k
		for {
			if b <= len(num) {
				x := num[0:a]
				y := num[b:]

				xy := fmt.Sprintf("%v%v", x, y)
				yz, _ := strconv.Atoi(xy)
				r = append(r, yz)
			} else {
				break
			}
			a = a + 1
			b = a + k
		}

		return r
	}
	rs := split(num, k)
	return fmt.Sprintf("%v", min(rs))
}

func TestRemoveKDigits(t *testing.T) {
	n := "1432219"
	k := 3
	r := removeKDigits(n, k)
	if r != "1219" {
		t.Errorf("Actual:%v, Expected:%v", r, "1219")
	}
	n = "10200"
	k = 1
	r = removeKDigits(n, k)
	if r != "200" {
		t.Errorf("Actual:%v, Expected:%v", r, "200")
	}

	n = "10"
	k = 2
	r = removeKDigits(n, k)
	if r != "0" {
		t.Errorf("Actual:%v, Expected:%v", r, "0")
	}

	n = "10001"
	k = 4
	r = removeKDigits(n, k)
	if r != "1" {
		t.Errorf("Actual:%v, Expected:%v", r, "0")
	}

}
