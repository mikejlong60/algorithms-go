package chapter4

import (
	"fmt"
	"github.com/mikejlong60/algorithms-go/chapter5"
)

type Process struct {
	id         int
	length     int
	deadline   int
	finishTime int
}

func (p Process) String() string {
	return fmt.Sprintf("Process{Id:%v, length:%v, deadline: %v, finishTime: %v, }", p.id, p.length, p.deadline, p.finishTime)
}

func MinimizeLateness(r []*Process) ([]*Process, *Process) {
	lt := func(l, r *Process) bool {
		if l.deadline < r.deadline {
			return true
		} else {
			return false
		}
	}
	rr := chapter5.MergeSort(r, lt)

	var timeline = 0
	var maxLate = &Process{}
	for _, j := range rr {
		timeline = timeline + j.length
		j.finishTime = timeline
		if maxLate.deadline < j.deadline {
			maxLate = j
		}
	}
	return rr, maxLate
}
