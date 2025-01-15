package chapter4

import (
	"github.com/mikejlong60/algorithms-go/chapter5"
)

type TimeSlot struct {
	id    int
	begin int
	end   int
}

func Schedule(r []*TimeSlot) []*TimeSlot {
	lt := func(l, r *TimeSlot) bool {
		if l.end < r.end {
			return true
		} else {
			return false
		}
	}
	rr := chapter5.MergeSort(r, lt)
	_, a := schedule(rr, []*TimeSlot{})
	return a
}

var totalSteps0 int

func schedule(r, a []*TimeSlot) ([]*TimeSlot, []*TimeSlot) {
	totalSteps0 = totalSteps0 + 1
	removeTimesThatStartBeforeXFinishes := func(x *TimeSlot, r []*TimeSlot) []*TimeSlot {
		var newR = []*TimeSlot{}
		for _, b := range r {
			if b.begin > x.end { //exclude b Timeslot because it overlaps with x
				newR = append(newR, b)
			}
			totalSteps0 = totalSteps0 + 1
		}
		return newR
	}

	if len(r) == 0 {
		return r, a
	} else {
		x := r[0]
		newR := removeTimesThatStartBeforeXFinishes(x, r)
		a = append(a, x)
		return schedule(newR, a)
	}
}
