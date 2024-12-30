package chapter4

import (
	"github.com/mikejlong60/algorithms/chapter5"
)

func ScheduleAll(r []*TimeSlot) [][]*TimeSlot { //Each row of the returned array is the schedule for a single resource(say a thread)
	lt := func(l, r *TimeSlot) bool {
		if l.begin < r.begin {
			return true
		} else {
			return false
		}
	}
	rr := chapter5.MergeSort(r, lt)
	_, a := scheduleAll(rr, [][]*TimeSlot{})
	return a
}

var totalSteps1 int

func scheduleAll(remainingTimeSlots []*TimeSlot, scheduledThreads [][]*TimeSlot) ([]*TimeSlot, [][]*TimeSlot) {
	timeSlotsOverlap := func(x *TimeSlot, y *TimeSlot) bool {
		if x.end > y.begin {
			return true
		} else {
			return false //exclude b Timeslot because it overlaps with x
		}
	}

	if len(remainingTimeSlots) == 0 {
		return remainingTimeSlots, scheduledThreads
	} else {
		//Iterate over remaining time slots
		nextThread := []*TimeSlot{}
		nextRemainingTimeSlots := []*TimeSlot{}
		for _, j := range remainingTimeSlots {
			totalSteps1 = totalSteps1 + 1
			if len(nextThread) == 0 {
				nextThread = append(nextThread, j)
			} else if !timeSlotsOverlap(nextThread[len(nextThread)-1], j) { // Add x to scheduled thread
				nextThread = append(nextThread, j)
			} else {
				nextRemainingTimeSlots = append(nextRemainingTimeSlots, j)
			}
		}
		scheduledThreads = append(scheduledThreads, nextThread)
		return scheduleAll(nextRemainingTimeSlots, scheduledThreads)
	}
}
