package chapter4

import (
	"fmt"
	"github.com/greymatter-io/golangz/arrays"
	"testing"
)

func TestAllIntervalScheduling(t *testing.T) {
	a1 := TimeSlot{
		id:    1,
		begin: 0,
		end:   3,
	}
	b2 := TimeSlot{
		id:    2,
		begin: 0,
		end:   8,
	}
	c3 := TimeSlot{
		id:    3,
		begin: 0,
		end:   3,
	}
	d4 := TimeSlot{
		id:    4,
		begin: 5,
		end:   8,
	}
	e5 := TimeSlot{
		id:    5,
		begin: 5,
		end:   14,
	}
	f6 := TimeSlot{
		id:    6,
		begin: 10,
		end:   15,
	}
	g7 := TimeSlot{
		id:    7,
		begin: 10,
		end:   15,
	}
	h8 := TimeSlot{
		id:    8,
		begin: 14,
		end:   20,
	}
	i9 := TimeSlot{
		id:    9,
		begin: 16,
		end:   20,
	}
	j10 := TimeSlot{
		id:    10,
		begin: 16,
		end:   20,
	}

	eq := func(l, r *TimeSlot) bool {
		if l.begin == r.begin && l.end == r.end && l.id == r.id {
			return true
		} else {
			return false
		}
	}
	r := []*TimeSlot{&a1, &b2, &c3, &d4, &e5, &f6, &g7, &h8, &i9, &j10}
	actual := ScheduleAll(r)
	expected := [][]*TimeSlot{{&c3, &e5, &h8}, {&b2, &g7, &j10}, {&a1, &d4, &f6, &i9}}
	for i, _ := range actual {
		if !arrays.ArrayEquality(actual[i], expected[i], eq) {
			t.Errorf("Actual:%v Expected:%v", actual, expected)
		}
	}
	fmt.Printf("total steps:%v", totalSteps1)
}
