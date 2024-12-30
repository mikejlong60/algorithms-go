package chapter4

import (
	"fmt"
	"github.com/greymatter-io/golangz/arrays"
	"testing"
)

func TestIntervalScheduling(t *testing.T) {
	a1 := TimeSlot{
		id:    1,
		begin: 0,
		end:   3,
	}
	a2 := TimeSlot{
		id:    2,
		begin: 0,
		end:   5,
	}
	a3 := TimeSlot{
		id:    3,
		begin: 4,
		end:   8,
	}
	a4 := TimeSlot{
		id:    4,
		begin: 6,
		end:   9,
	}
	a5 := TimeSlot{
		id:    5,
		begin: 10,
		end:   15,
	}
	a6 := TimeSlot{
		id:    6,
		begin: 0,
		end:   16,
	}
	a7 := TimeSlot{
		id:    7,
		begin: 11,
		end:   17,
	}
	a8 := TimeSlot{
		id:    8,
		begin: 19,
		end:   22,
	}
	a9 := TimeSlot{
		id:    9,
		begin: 18,
		end:   24,
	}

	eq := func(l, r *TimeSlot) bool {
		if l.begin == r.begin && l.end == r.end {
			return true
		} else {
			return false
		}
	}
	r := []*TimeSlot{&a1, &a2, &a3, &a4, &a5, &a6, &a7, &a8, &a9}
	actual := Schedule(r)
	expected := []*TimeSlot{&a1, &a3, &a5, &a8}
	if !arrays.ArrayEquality(actual, expected, eq) {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
	fmt.Printf("totalSteps:%v", totalSteps0)
}
