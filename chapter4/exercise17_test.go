package chapter4

import (
	"github.com/greymatter-io/golangz/sets"
	"testing"
)

func TestExercise17(t *testing.T) {
	//This is the example from the book.  The book gives an answer of j1 and j3
	//being the optimal solution. But there are others as well. This example
	//shows another valid schedule: j2 and j1. The algorithm I used was based upon
	//the greedy rule that chooses the earliest finisher and eliminates the
	//conflicting jobs, and then chooses the next earliest finisher.
	//Some caveats:
	//1. I have adjusted the beginning and ending times from the AM/PM times in the book to use a 24 hour clock.
	//2. In the event that a Timeslot goes into the next day I added 24 hours to the ending time.  I reasoned
	//   this is the way Airlines determine start and finish times for a flight.  And also reasoned that
	//   the task of creating a schedule must occur from a starting point, say 0:00 on the 24 hr clock. It makes
	//   no sense to schedule a job as starting in the past.  This is a monotonic clock.
	j0 := TimeSlot{
		id:    0,
		begin: 18,
		end:   30,
	}
	j1 := TimeSlot{
		id:    1,
		begin: 21,
		end:   28,
	}
	j2 := TimeSlot{
		id:    2,
		begin: 3,
		end:   14,
	}
	j3 := TimeSlot{
		id:    3,
		begin: 13,
		end:   19,
	}
	eq := func(l, r *TimeSlot) bool {
		if l.begin == r.begin && l.end == r.end && l.id == r.id {
			return true
		} else {
			return false
		}
	}
	r := []*TimeSlot{&j0, &j1, &j2, &j3}
	actual := Schedule(r)
	expected := []*TimeSlot{&j1, &j2}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
}

func TestExercise17b(t *testing.T) {
	//This is a slight variation from the book, a new job that can fit in between the
	//optimal solution, j4. The algorithm correctly adds the new job

	j0 := TimeSlot{
		id:    0,
		begin: 18,
		end:   30,
	}
	j1 := TimeSlot{
		id:    1,
		begin: 21,
		end:   28,
	}
	j2 := TimeSlot{
		id:    2,
		begin: 3,
		end:   14,
	}
	j3 := TimeSlot{
		id:    3,
		begin: 13,
		end:   19,
	}
	j4 := TimeSlot{
		id:    4,
		begin: 15,
		end:   16,
	}
	eq := func(l, r *TimeSlot) bool {
		if l.begin == r.begin && l.end == r.end && l.id == r.id {
			return true
		} else {
			return false
		}
	}
	r := []*TimeSlot{&j0, &j1, &j2, &j3, &j4}
	actual := Schedule(r)
	expected := []*TimeSlot{&j1, &j2, &j4}
	if !sets.SetEquality(actual, expected, eq) {
		t.Errorf("Actual:%v Expected:%v", actual, expected)
	}
}
