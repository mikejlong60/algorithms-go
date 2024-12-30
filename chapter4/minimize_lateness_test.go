package chapter4

import (
	"github.com/greymatter-io/golangz/arrays"
	"testing"
)

func TestScheduleWithDeadline1(t *testing.T) {
	job1 := Process{
		id:       1,
		length:   1,
		deadline: 2,
	}
	job2 := Process{
		id:       2,
		length:   2,
		deadline: 4,
	}

	job3 := Process{
		id:       2,
		length:   3,
		deadline: 6,
	}

	eq := func(l, r *Process) bool {
		if l.id == r.id {
			return true
		} else {
			return false
		}
	}
	r := []*Process{&job1, &job2, &job3}
	actual, maxLate := MinimizeLateness(r)
	expected := []*Process{&job1, &job2, &job3}
	if !(arrays.ArrayEquality(actual, expected, eq) && maxLate == &job3 && maxLate.finishTime == 6 && maxLate.finishTime-maxLate.deadline <= 0) {
		t.Errorf("Actual Schedule:%v Expected Schedule:%v, Max Late:=%v", actual, expected, maxLate)
	}
	//if !arrays.ArrayEquality(actual, expected, eq) && maxLate != &job2 && maxLate.finishTime != 11 {
	//	t.Errorf("Actual Schedule:%v Expected Schedule:%v, Max Late:=%v", actual, expected, maxLate)
	//}
}

func TestScheduleWithDeadline2(t *testing.T) {
	job1 := Process{
		id:       1,
		length:   1,
		deadline: 100,
	}
	job2 := Process{
		id:       2,
		length:   10,
		deadline: 10,
	}

	eq := func(l, r *Process) bool {
		if l.id == r.id {
			return true
		} else {
			return false
		}
	}
	r := []*Process{&job1, &job2}
	actual, maxLate := MinimizeLateness(r)
	expected := []*Process{&job2, &job1}
	if !(arrays.ArrayEquality(actual, expected, eq) && maxLate == &job1 && maxLate.finishTime == 11 && maxLate.finishTime-maxLate.deadline <= 0) {
		t.Errorf("Actual Schedule:%v Expected Schedule:%v, Max Late:=%v", actual, expected, maxLate)
	}
}

func TestScheduleWithDeadline3(t *testing.T) {
	job1 := Process{
		id:       1,
		length:   1,
		deadline: 2,
	}
	job2 := Process{
		id:       2,
		length:   10,
		deadline: 10,
	}

	eq := func(l, r *Process) bool {
		if l.id == r.id {
			return true
		} else {
			return false
		}
	}
	r := []*Process{&job1, &job2}
	actual, maxLate := MinimizeLateness(r)
	expected := []*Process{&job1, &job2}
	if !(arrays.ArrayEquality(actual, expected, eq) && maxLate == &job2 && maxLate.finishTime == 11 && maxLate.finishTime-maxLate.deadline == 1) {
		t.Errorf("Actual Schedule:%v Expected Schedule:%v, Max Late:=%v", actual, expected, maxLate)
	}
}
